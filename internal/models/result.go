package models

import "time"

// Result struct of Data for Epever (Epsolar) Controller LS-B series
type ResultLSB struct {
	Connected bool        `json:"connected"`
	Updated   time.Time   `json:"updated"`
	Interval  int         `json:"update_interval,omitempty"`
	Model     string      `json:"model,omitempty"`
	Device    *SummaryLSB `json:"device,omitempty"`
}
