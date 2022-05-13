package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Device struct {
		Port           string
		UpdateInterval int
		Name           string
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

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, ErrConfigFileRead.Error())
	}

	cnf := &Config{}
	if err := viper.Unmarshal(cnf); err != nil {
		return nil, errors.Wrap(err, ErrConfigUnmarshal.Error())
	}

	return cnf, nil
}
