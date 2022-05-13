package models

import "time"

type ResultLSX024B struct {
	Connected string          `json:"connected"`
	Updated   time.Time       `json:"updated"`
	Interval  int             `json:"update_interval,omitempty"`
	Device    *SummaryLSX024B `json:"device,omitempty"`
}
