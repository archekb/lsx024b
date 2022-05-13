package mqtt

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type MQTT struct {
	client   mqtt.Client
	topic    string
	haDevice *haDevice
}

func New(addr, user, password, clientId string) MQTT {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(addr)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	clientId = strings.ReplaceAll(strings.ToLower(clientId), " ", "_")
	opts.SetClientID(clientId)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(2 * time.Minute)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("MQTT default publish handler: [%s] -> [%s]", msg.Topic(), string(msg.Payload()))
	})

	return MQTT{client: mqtt.NewClient(opts)}
}

func (c *MQTT) Connect() error {
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (c *MQTT) IsConnected() bool {
	if c.client == nil {
		return false
	}

	return c.client.IsConnected()
}

func (c *MQTT) Close() {
	if !c.IsConnected() {
		return
	}

	c.client.Disconnect(0)
}

func (c *MQTT) SetTopic(topic, name string) {
	smallName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	topicSplited := TopicPrepare(topic)
	topicSplited = append(topicSplited, smallName)
	c.topic = strings.Join(topicSplited, "/")
}

func (c *MQTT) PublishToDefault(message interface{}) error {
	return c.Publish(c.topic, message)
}

func (c *MQTT) Publish(topic string, message interface{}) error {
	if !c.IsConnected() {
		return ErrNotConnected
	}

	if topic == "" {
		return ErrTopicIsEmpty
	}

	m, _ := json.Marshal(message)

	t := c.client.Publish(topic, 0, false, string(m))
	t.Wait()
	if t.Error() != nil {
		log.Println(t.Error())
		return t.Error()
	}

	return nil
}
