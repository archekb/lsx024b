package main

import (
	"embed"
	"lsx024b/internal/config"
	"lsx024b/internal/dev"
	"lsx024b/internal/http"
	"lsx024b/internal/mqtt"
	"strings"
	"time"

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

	smallName := strings.ReplaceAll(strings.ToLower(cnf.Device.Name), " ", "_")
	topicA := mqtt.TopicPrepare(cnf.MQTT.Topic)
	topicA = append(topicA, smallName)
	topic := strings.Join(topicA, "/")

	d, err := dev.New(cnf.Device.Port)
	if err != nil {
		log.Fatalln(err)
	}

	var mqttClient *mqtt.MQTT
	var hd *mqtt.HADevice
	if cnf.MQTT.Address != "" {
		cr, err := d.ControllerRated()
		if err != nil {
			log.Fatalln(err)
		}

		mqttClient, _ = mqtt.New(cnf.MQTT.Address, cnf.MQTT.User, cnf.MQTT.Password, smallName)

		if cnf.MQTT.HomeAssistant {
			hd = mqtt.NewHADevice(cnf.Device.Name, int(cr.OutputCurrentOfLoad), topic)

			mqttClient.Publish(hd.GenerateSensor("Input Voltage", "controller_real_time.input_voltage", "voltage", "V"))
			mqttClient.Publish(hd.GenerateSensor("Input Power", "controller_real_time.input_power", "power", "W"))
			mqttClient.Publish(hd.GenerateSensor("Input Current", "controller_real_time.input_current", "current", "A"))

			mqttClient.Publish(hd.GenerateSensor("Output Voltage", "controller_real_time.output_voltage", "voltage", "V"))
			mqttClient.Publish(hd.GenerateSensor("Output Power", "controller_real_time.output_power", "power", "W"))
			mqttClient.Publish(hd.GenerateSensor("Output Current", "controller_real_time.output_current", "current", "A"))

			mqttClient.Publish(hd.GenerateSensor("Battery", "controller_real_time.battery_soc", "battery", "%"))
			mqttClient.Publish(hd.GenerateSensor("Battery Status", "controller_status.battery", "", ""))
			mqttClient.Publish(hd.GenerateSensor("Battery Voltage", "controller_statistical.battery_voltage", "voltage", "V"))
			mqttClient.Publish(hd.GenerateSensor("Battery Current", "controller_statistical.battery_current", "current", "A"))
			mqttClient.Publish(hd.GenerateSensor("Battery Temperature", "controller_real_time.battery_temperature", "temperature", "℃"))

			mqttClient.Publish(hd.GenerateSensor("Equipment Temperature", "controller_real_time.equipment_inside_temperature", "temperature", "℃"))
			mqttClient.Publish(hd.GenerateSensor("Charging", "controller_status.charging_type", "", ""))
		}
	}

	type currentState struct {
		Status   string    `json:"status"`
		Updated  time.Time `json:"updated"`
		Interval int       `json:"update_interval"`
		*dev.ControllerSummary
	}

	var lastState *currentState
	stopReadController := make(chan struct{}, 1)
	prevOutputSource := "output"
	go func() {
		ticker := time.Tick(time.Duration(cnf.Device.UpdateInterval) * time.Second)
		for {
			select {
			case <-stopReadController:
				return
			case <-ticker:
				dataControllerSummary, err := d.ControllerSummary()
				if err != nil {
					lastState = &currentState{Status: "offline", Updated: time.Now().UTC()}
					log.Println(err)
					if mqttClient != nil {
						mqttClient.Publish(topic, lastState)
					}
					continue
				}
				lastState = &currentState{Status: "online", Updated: time.Now().UTC(), Interval: cnf.Device.UpdateInterval, ControllerSummary: dataControllerSummary}

				if mqttClient != nil {
					if dataControllerSummary.ControllerRealTime.OutputCurrent == 0 && dataControllerSummary.ControllerRealTime.LoadCurrent > 0 {
						if prevOutputSource != "load" {
							prevOutputSource = "load"
							mqttClient.Publish(hd.GenerateSensor("Output Power", "controller_real_time.load_power", "power", "W"))
							mqttClient.Publish(hd.GenerateSensor("Output Current", "controller_real_time.load_current", "current", "A"))
							mqttClient.Publish(hd.GenerateSensor("Output Voltage", "controller_real_time.load_voltage", "voltage", "V"))
						}
					} else {
						if prevOutputSource != "output" {
							prevOutputSource = "output"
							mqttClient.Publish(hd.GenerateSensor("Output Power", "controller_real_time.output_power", "power", "W"))
							mqttClient.Publish(hd.GenerateSensor("Output Current", "controller_real_time.output_current", "current", "A"))
							mqttClient.Publish(hd.GenerateSensor("Output Voltage", "controller_real_time.output_voltage", "voltage", "V"))
						}
					}
					mqttClient.Publish(topic, lastState)
				}
			}
		}
	}()

	// server HTTP
	srvHTTP, err := http.New(&staticFS, &http.Env{
		State: &lastState,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if cnf.HTTP.Certificate != "" && cnf.HTTP.Key != "" {
		srvHTTP.RunTLS(cnf.HTTP.Address, cnf.HTTP.Certificate, cnf.HTTP.Key)
	} else {
		srvHTTP.Run(cnf.HTTP.Address)
	}

	// garceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	stopReadController <- struct{}{}

	if mqttClient != nil {
		mqttClient.Publish(topic, currentState{Status: "offline", Updated: time.Now().UTC()})
		mqttClient.Close()
	}

	d.Close()

	log.Infoln("Server exiting")
}
