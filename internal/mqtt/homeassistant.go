package mqtt

import (
	"fmt"
	"strings"
)

type HADeviceDescription struct {
	Identifiers  []string `json:"identifiers"`
	Manufacturer string   `json:"manufacturer"`
	Model        string   `json:"model"`
	Name         string   `json:"name"`
	SW           string   `json:"sw_version"`
}

type HAAvailabilityDescription struct {
	PayloadAvailable     string `json:"payload_available,omitempty"`
	PayloadNotAvailable  string `json:"payload_not_available,omitempty"`
	AvailabilityTemplate string `json:"availability_template,omitempty"`
	AvailabilityTopic    string `json:"availability_topic,omitempty"`
}

type HASensorDescription struct {
	Device *HADeviceDescription `json:"device"`
	HAAvailabilityDescription

	DeviceClass         string `json:"device_class,omitempty"`
	EnabledByDefault    bool   `json:"enabled_by_default"`
	EntityCategory      string `json:"entity_category,omitempty"`
	Name                string `json:"name"`
	StateTopic          string `json:"state_topic"`
	UniqueId            string `json:"unique_id"`
	UnitOfMeasurement   string `json:"unit_of_measurement,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	ValueTemplate       string `json:"value_template,omitempty"`
}

type HADevice struct {
	topic        []string
	device       *HADeviceDescription
	availability *HAAvailabilityDescription
}

func NewHADevice(name string, model int, topic string) *HADevice {
	// nameSanitized := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	topicSplited := TopicPrepare(topic)

	had := &HADevice{
		topic: topicSplited,
		device: &HADeviceDescription{
			Identifiers:  []string{strings.Join(topicSplited, "_")},
			Manufacturer: "Epsolar",
			Model:        fmt.Sprintf("LS%d24B", model),
			Name:         name,
			SW:           "0.0.1",
		},
		availability: &HAAvailabilityDescription{
			AvailabilityTopic:    strings.Join(topicSplited, "/"),
			AvailabilityTemplate: "{{ value_json.status }}",
		},
	}

	return had
}

func (had *HADevice) DeviceUID() string {
	return strings.Join(had.topic, "_")
}

func (had *HADevice) GenerateSensor(name, jpath, class, unit string) (string, *HASensorDescription) {
	smallName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	haTopic := fmt.Sprintf("homeassistant/sensor/%s/%s/config", had.DeviceUID(), smallName)

	return haTopic, &HASensorDescription{
		Device:                    had.device,
		HAAvailabilityDescription: *had.availability,

		DeviceClass:       class,
		EnabledByDefault:  true,
		EntityCategory:    "diagnostic",
		Name:              fmt.Sprintf("%s %s", had.device.Name, name),
		StateTopic:        strings.Join(had.topic, "/"),
		UniqueId:          strings.Join(append(had.topic, smallName), "_"),
		UnitOfMeasurement: unit,
		// JSONAttributesTopic: topic,
		ValueTemplate: fmt.Sprintf("{{ value_json.%s }}", jpath),
	}
}

func TopicPrepare(topic string) []string {
	topicSplited := strings.Split(topic, "/")
	newTopicSplited := []string{}

	for i, v := range topicSplited {
		if topicSplited[i] == "" {
			continue
		}
		newTopicSplited = append(newTopicSplited, v)
	}

	return newTopicSplited
}
