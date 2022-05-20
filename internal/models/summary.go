package models

// Summary struct of Data for Epever (Epsolar) Controller LS-B series
type SummaryLSB struct {
	Rated       *LSB_Rated       `json:"rated,omitempty"`
	RealTime    *LSB_RealTime    `json:"real_time,omitempty"`
	Status      *LSB_Status      `json:"status,omitempty"`
	Statistical *LSB_Statistical `json:"statistical,omitempty"`
	Settings    *LSB_Settings    `json:"settings,omitempty"`
	Switches    *LSB_Switches    `json:"switches,omitempty"`
	Discrete    *LSB_Discrete    `json:"discrete,omitempty"`
}
