package models

type SummaryLSX024B struct {
	LSX024BRated       *LSX024BRated       `json:"rated,omitempty"`
	LSX024BRealTime    *LSX024BRealTime    `json:"real_time,omitempty"`
	LSX024BStatus      *LSX024BStatus      `json:"status,omitempty"`
	LSX024BStatistical *LSX024BStatistical `json:"statistical,omitempty"`
	LSX024BSettings    *LSX024BSettings    `json:"settings,omitempty"`
	LSX024BSwitches    *LSX024BSwitches    `json:"switches,omitempty"`
	LSX024BDiscrete    *LSX024BDiscrete    `json:"discrete,omitempty"`
}
