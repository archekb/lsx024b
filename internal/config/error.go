package config

import "errors"

var (
	ErrConfigFileRead                 = errors.New("config file read error")
	ErrConfigUnmarshal                = errors.New("config unmarshal error")
	ErrDevicePortIsEmpty              = errors.New("device Port is empty")
	ErrDeviceSlaveIdIsWrong           = errors.New("device SlaveId is wrong")
	ErrDeviceNameIsEmpty              = errors.New("device Name is empty")
	ErrDeviceUpdateIntervalIsNegative = errors.New("device UpdateInterval is negative")
	ErrHTTPServerKeyOrCertIsEmpty     = errors.New("HTTP server Key or Certificate is empty")
	ErrHTTPAddressFormat              = errors.New("HTTP server Address should be 'address:port'")
	ErrMQTTAddressFormat              = errors.New("MQTT server Address should be 'scheme://address:port', scheme can be 'tcp', 'ssl' or 'ws'")
)
