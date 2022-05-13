package mqtt

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type MQTT struct {
	client mqtt.Client
}

func New(addr, user, password, clientId string) (*MQTT, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(addr)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetClientID(clientId)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(2 * time.Minute)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("MQTT default publish handler: [%s] -> [%s]", msg.Topic(), string(msg.Payload()))
	})

	mc := mqtt.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MQTT{client: mc}, nil
}

func (c *MQTT) IsConnected() bool {
	return c.client.IsConnected()
}

func (c *MQTT) Close() {
	c.client.Disconnect(0)
}

func (c *MQTT) Publish(topic string, message interface{}) {
	m, _ := json.Marshal(message)

	t := c.client.Publish(topic, 0, false, string(m))
	t.Wait()
	if t.Error() != nil {
		log.Println(t.Error())
	}
}
