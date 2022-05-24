package mqtt

import "errors"

var (
	ErrNotConnected = errors.New("not connected")
	ErrTopicIsEmpty = errors.New("topic is empty")
	ErrHANotInit    = errors.New("Home Assissstant not init")
)
