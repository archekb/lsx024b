package mqtt

import (
	"fmt"
	"strings"
)

const HA_DEFAULT_TOPIC = "homeassistant"

func (c *MQTT) HAInit(name, model, version string) {
	identifier := strings.Join(TopicPrepare(c.topic), "_")

	c.haDevice = &haDevice{
		topic:            TopicPrepare(c.topic),
		prevOutputSource: "output",
		device: &haDeviceDescription{
			Identifiers:  []string{identifier},
			Manufacturer: "Epever (Epsolar)",
			Model:        model,
			Name:         name,
			SW:           version,
		},
		availability: &haAvailabilityDescription{
			AvailabilityTopic:    c.topic,
			AvailabilityTemplate: "{{ value_json.connected }}",
		},
	}
}

func (c *MQTT) HAPublishDevice() error {
	if !c.IsConnected() {
		return ErrNotConnected
	}

	if c.haDevice == nil {
		return ErrHANotInit
	}

	c.Publish(c.haDevice.GenerateSensor("Input Voltage", "device.real_time.input_voltage", "voltage", "V"))
	c.Publish(c.haDevice.GenerateSensor("Input Power", "device.real_time.input_power", "power", "W"))
	c.Publish(c.haDevice.GenerateSensor("Input Current", "device.real_time.input_current", "current", "A"))

	c.Publish(c.haDevice.GenerateSensor("Output Voltage", "device.real_time.output_voltage", "voltage", "V"))
	c.Publish(c.haDevice.GenerateSensor("Output Power", "device.real_time.output_power", "power", "W"))
	c.Publish(c.haDevice.GenerateSensor("Output Current", "device.real_time.output_current", "current", "A"))

	c.Publish(c.haDevice.GenerateSensor("Load Voltage", "device.real_time.load_voltage", "voltage", "V"))
	c.Publish(c.haDevice.GenerateSensor("Load Power", "device.real_time.load_power", "power", "W"))
	c.Publish(c.haDevice.GenerateSensor("Load Current", "device.real_time.load_current", "current", "A"))

	c.Publish(c.haDevice.GenerateSensor("Battery", "device.real_time.battery_soc", "battery", "%"))
	c.Publish(c.haDevice.GenerateSensor("Battery Status", "device.status.battery", "", ""))
	c.Publish(c.haDevice.GenerateSensor("Battery Voltage", "device.statistical.battery_voltage", "voltage", "V"))
	c.Publish(c.haDevice.GenerateSensor("Battery Current", "device.statistical.battery_current", "current", "A"))
	c.Publish(c.haDevice.GenerateSensor("Battery Temperature", "device.real_time.battery_temperature", "temperature", "℃"))

	c.Publish(c.haDevice.GenerateSensor("Equipment Temperature", "device.real_time.equipment_inside_temperature", "temperature", "℃"))
	c.Publish(c.haDevice.GenerateSensor("Charging", "device.status.charging_type", "", ""))

	return nil
}

type haDeviceDescription struct {
	Identifiers  []string `json:"identifiers"`
	Manufacturer string   `json:"manufacturer"`
	Model        string   `json:"model"`
	Name         string   `json:"name"`
	SW           string   `json:"sw_version"`
}

type haAvailabilityDescription struct {
	PayloadAvailable     string `json:"payload_available,omitempty"`
	PayloadNotAvailable  string `json:"payload_not_available,omitempty"`
	AvailabilityTemplate string `json:"availability_template,omitempty"`
	AvailabilityTopic    string `json:"availability_topic,omitempty"`
}

type haSensorDescription struct {
	Device *haDeviceDescription `json:"device"`
	haAvailabilityDescription

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

type haDevice struct {
	topic            []string
	prevOutputSource string
	device           *haDeviceDescription
	availability     *haAvailabilityDescription
}

func (had *haDevice) DeviceUID() string {
	return strings.ReplaceAll(strings.ToLower(had.device.Name), "_", "_")
}

func (had *haDevice) GenerateSensor(name, jpath, class, unit string) (string, *haSensorDescription) {
	smallName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	haTopic := fmt.Sprintf("%s/sensor/%s/%s/config", HA_DEFAULT_TOPIC, had.DeviceUID(), smallName)

	return haTopic, &haSensorDescription{
		Device:                    had.device,
		haAvailabilityDescription: *had.availability,

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
