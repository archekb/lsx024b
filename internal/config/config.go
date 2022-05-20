package config

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Device struct {
		Port           string
		Name           string
		Model          string
		UpdateInterval int
		SlaveId        int
	}

	HTTP struct {
		Address     string
		Certificate string
		Key         string
	}

	MQTT struct {
		Address       string
		User          string
		Password      string
		Topic         string
		HomeAssistant bool
	}
}

func New(filename string) (*Config, error) {
	viper.SetConfigFile(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("Device.UpdateInterval", 10)
	viper.SetDefault("Device.SlaveId", 1)
	viper.SetDefault("Device.Model", "LS-B Compatible")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, ErrConfigFileRead.Error())
	}

	cnf := &Config{}
	if err := viper.Unmarshal(cnf); err != nil {
		return nil, errors.Wrap(err, ErrConfigUnmarshal.Error())
	}

	if err := Validate(cnf); err != nil {
		return nil, err
	}

	return cnf, nil
}

func Validate(cnf *Config) error {
	if cnf.Device.Port == "" {
		return ErrDevicePortIsEmpty
	}

	if cnf.Device.SlaveId <= 0 || cnf.Device.SlaveId > 247 {
		return ErrDeviceSlaveIdIsWrong
	}

	if cnf.Device.Name == "" {
		return ErrDeviceNameIsEmpty
	}

	if cnf.Device.UpdateInterval < 0 {
		return ErrDeviceUpdateIntervalIsNegative
	}

	if cnf.HTTP.Address != "" && !strings.Contains(cnf.HTTP.Address, ":") {
		return ErrHTTPAddressFormat
	}

	if (cnf.HTTP.Certificate != "" && cnf.HTTP.Key == "") || (cnf.HTTP.Certificate == "" && cnf.HTTP.Key != "") {
		return ErrHTTPServerKeyOrCertIsEmpty
	}

	re, _ := regexp.Compile(`(tcp|ws|ssl):\/\/.*:\d+`)

	if cnf.MQTT.Address != "" && !re.MatchString(cnf.MQTT.Address) {
		return ErrMQTTAddressFormat
	}

	return nil
}
