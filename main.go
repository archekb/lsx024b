package main

import (
	"embed"
	"time"

	"github.com/archekb/lsx024b/internal/config"
	"github.com/archekb/lsx024b/internal/device"
	"github.com/archekb/lsx024b/internal/http"
	"github.com/archekb/lsx024b/internal/models"
	"github.com/archekb/lsx024b/internal/mqtt"

	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

//go:embed web/*
var staticFS embed.FS

func init() {
	// logger
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "15-01-2006 15:04:05.000000", FullTimestamp: true, ForceColors: true})
	log.SetLevel(log.TraceLevel)
}

func main() {
	cnf, err := config.New("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	dev := device.NewLSX024B(cnf.Device.Port, cnf.Device.SlaveId)
	if dev.Connect() != nil {
		log.Fatalln(err)
	}

	var mqttClient mqtt.MQTT
	if cnf.MQTT.Address != "" {
		mqttClient = mqtt.New(cnf.MQTT.Address, cnf.MQTT.User, cnf.MQTT.Password, cnf.Device.Name)
		mqttClient.SetTopic(cnf.MQTT.Topic, cnf.Device.Name)
		mqttClient.Connect()

		if cnf.MQTT.HomeAssistant {
			cr, err := dev.GetRated()
			if err != nil {
				log.Errorln(err)
			}

			mqttClient.HAInit(cnf.Device.Name, int(cr.OutputCurrentOfLoad))
			mqttClient.HAPublishDevice()
		}
	}

	var lastRes *models.ResultLSX024B = &models.ResultLSX024B{}
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
					lastRes = &models.ResultLSX024B{Connected: "offline", Updated: time.Now().UTC()}
					log.Println(err)
					mqttClient.PublishToDefault(lastRes)
					continue
				}
				lastRes = &models.ResultLSX024B{Connected: "online", Updated: time.Now().UTC(), Interval: cnf.Device.UpdateInterval, Device: dataSummary}

				mqttClient.HAPublishUpdateDevice(dataSummary.LSX024BRealTime.OutputCurrent, dataSummary.LSX024BRealTime.LoadCurrent)
				mqttClient.PublishToDefault(lastRes)
			}
		}
	}()

	// server HTTP
	if cnf.HTTP.Address != "" {
		srvHTTP, err := http.New(&staticFS, &http.Env{
			State: &lastRes,
		})
		if err != nil {
			log.Fatalln(err)
		}

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

	log.Println("Shutting down server...")
	stopReadController <- struct{}{}

	mqttClient.PublishToDefault(&models.ResultLSX024B{Connected: "offline", Updated: time.Now().UTC()})
	mqttClient.Close()

	dev.Close()

	log.Infoln("Server exiting")
}
