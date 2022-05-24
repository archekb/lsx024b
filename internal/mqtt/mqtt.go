package mqtt

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/archekb/lsx024b/internal/log"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	client    mqtt.Client
	topic     string
	haDevice  *haDevice
	onConnect func()
}

func New(addr, user, password, clientId string) *MQTT {
	m := MQTT{}

	opts := mqtt.NewClientOptions()

	opts.AddBroker(addr)
	opts.SetUsername(user)
	opts.SetPassword(password)

	opts.SetKeepAlive(5 * time.Second)
	opts.SetPingTimeout(5 * time.Second)

	clientId = fmt.Sprintf("%s_%s", strings.ReplaceAll(strings.ToLower(clientId), " ", "_"), randSeq(16))
	opts.SetClientID(clientId)

	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(2 * time.Minute)

	opts.SetOnConnectHandler(func(c mqtt.Client) {
		if m.onConnect != nil {
			m.onConnect()
		}
	})

	mqtt.ERROR = log.StandartNamed("MQTT error")
	mqtt.CRITICAL = log.StandartNamed("MQTT critical")
	mqtt.WARN = log.StandartNamed("MQTT warning")
	// mqtt.DEBUG = log.StandartNamed("MQTT debug")

	m.client = mqtt.NewClient(opts)
	return &m
}

func (c *MQTT) Connect() error {
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	c.client.OptionsReader()
	return nil
}

func (c *MQTT) IsConnected() bool {
	if c.client == nil {
		return false
	}

	return c.client.IsConnectionOpen()
}

func (c *MQTT) Close() {
	if !c.IsConnected() {
		return
	}

	c.client.Disconnect(0)
}

func (c *MQTT) SetOnConnect(f func()) {
	c.onConnect = f
}

func (c *MQTT) SetDefaultTopic(topic, name string) {
	smallName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	topicSplited := TopicPrepare(topic)
	topicSplited = append(topicSplited, smallName)
	c.topic = strings.Join(topicSplited, "/")
}

func (c *MQTT) PublishToDefault(message interface{}) error {
	return c.Publish(c.topic, false, message)
}

func (c *MQTT) Publish(topic string, retained bool, message interface{}) error {
	if !c.IsConnected() {
		return ErrNotConnected
	}

	if topic == "" {
		return ErrTopicIsEmpty
	}

	m, _ := json.Marshal(message)

	t := c.client.Publish(topic, 0, retained, string(m))
	if t.Wait(); t.Error() != nil {
		log.Error("MQTT Publish error:", t.Error())
		return t.Error()
	}

	return nil
}
