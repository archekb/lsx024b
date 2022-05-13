package config

import "errors"

var (
	ErrConfigFileRead  = errors.New("error config file read")
	ErrConfigUnmarshal = errors.New("error config unmarshal")
)
