package main

import (
	"embed"
	"time"

	"github.com/archekb/lsx024b/internal/config"
	"github.com/archekb/lsx024b/internal/device"
	"github.com/archekb/lsx024b/internal/http"
	"github.com/archekb/lsx024b/internal/log"
	"github.com/archekb/lsx024b/internal/models"
	"github.com/archekb/lsx024b/internal/mqtt"

	"os"
	"os/signal"
	"syscall"
)

//go:embed web/*
var staticFS embed.FS

var version string

func init() {
	if version != "" {
		log.ProductionMode()
	}
}

func main() {
	log.Infof("Starting lsx024b v%s", version)

	cnf, err := config.New("config.yml")
	if err != nil {
		log.Fatal("Configuration error: ", err)
	}

	// connect to Device
	dev := device.NewLSB(cnf.Device.Port, cnf.Device.SlaveId)
	if err := dev.Connect(); err != nil {
		log.Fatalf("Can't connect to %s: %s", cnf.Device.Port, err)
	}
	log.Infof("Device via %s connected", cnf.Device.Port)

	// connect to MQTT
	var mqttClient mqtt.MQTT
	if cnf.MQTT.Address != "" {
		mqttClient = mqtt.New(cnf.MQTT.Address, cnf.MQTT.User, cnf.MQTT.Password, cnf.Device.Name)
		mqttClient.SetTopic(cnf.MQTT.Topic, cnf.Device.Name)

		if err := mqttClient.Connect(); err != nil {
			log.Fatal("Can't connect to MQTT: ", err)
		}
		log.Info("Connected to MQTT server")

		if cnf.MQTT.HomeAssistant {
			mqttClient.HAInit(cnf.Device.Name, cnf.Device.Model, version)
			mqttClient.HAPublishDevice()
		}
	}

	// read data from device
	var lastRes *models.ResultLSB
	stopReadController := make(chan struct{}, 1)
	go func() {
		ticker := time.Tick(time.Duration(cnf.Device.UpdateInterval) * time.Second)
		for {
			select {
			case <-stopReadController:
				return

			case <-ticker:
				dataSummary, err := dev.Summary()
				if err != nil {
					log.Error("Error reading data from controller: ", err)

					lastRes = &models.ResultLSB{Connected: "offline", Updated: time.Now().UTC(), Model: cnf.Device.Model}
					if mqttClient.IsConnected() {
						mqttClient.PublishToDefault(lastRes)
					}
					continue
				}

				lastRes = &models.ResultLSB{Connected: "online", Updated: time.Now().UTC(), Interval: cnf.Device.UpdateInterval, Model: cnf.Device.Model, Device: dataSummary}
				if mqttClient.IsConnected() {
					mqttClient.PublishToDefault(lastRes)
				}
			}
		}
	}()

	// server HTTP
	if cnf.HTTP.Address != "" {
		log.Info("Starting HTTP Server")

		srvHTTP := http.New(&staticFS, &http.Env{State: &lastRes, Release: version != ""})
		if cnf.HTTP.Certificate != "" && cnf.HTTP.Key != "" {
			srvHTTP.RunTLS(cnf.HTTP.Address, cnf.HTTP.Certificate, cnf.HTTP.Key)
		} else {
			srvHTTP.Run(cnf.HTTP.Address)
		}
	}

	// garceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")
	stopReadController <- struct{}{}

	if mqttClient.IsConnected() {
		mqttClient.PublishToDefault(&models.ResultLSB{Connected: "offline", Updated: time.Now().UTC()})
		mqttClient.Close()
	}
	dev.Close()

	log.Info("Server exiting")
}
