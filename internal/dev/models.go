package dev

// This section discribe rated (nominals) of controller (not real only nominals)
type ControllerRated struct {
	InputVoltage float64 `addr:"0x3000" divide:"100" json:"input_voltage"`              // [V] Charging equipment rated input voltage (PV array rated voltage)
	InputCurrent float64 `addr:"0x3001" divide:"100" json:"input_current"`              // [A] Charging equipment rated input current  (PV array rated current)
	InputPower   float64 `addr:"0x3003" laddr:"0x3002" divide:"100" json:"input_power"` // [W] Charging equipment rated input power (PV array rated power)

	OutputVoltage float64 `addr:"0x3004" divide:"100" json:"output_voltage"`              // [V] Charging equipment rated output voltage (Rated charging voltage to battery)
	OutputCurrent float64 `addr:"0x3005" divide:"100" json:"output_current"`              // [A] Charging equipment rated output current (Rated charging current to battery)
	OutputPower   float64 `addr:"0x3007" laddr:"0x3006" divide:"100" json:"output_power"` // [W] Charging equipment rated output power (Rated charging power to battery)

	ChargingMode        string  `addr:"0x3008" enum:"00:Connect/disconnect, 01:PWM, 02:MPP" json:"charging_mode"`
	OutputCurrentOfLoad float64 `addr:"0x300E" divide:"100" json:"output_current_of_load"` // [A] Rated output current of load (LS1024B - 10, LS2024B - 20, LS3024B - 30)
}

// This section discribe realtime status of controller
type ControllerRealTime struct {
	InputVoltage float64 `addr:"0x3100" divide:"100" json:"input_voltage"`              // [V] Charging equipment input voltage (Solar charge controller--PV array voltage)
	InputCurrent float64 `addr:"0x3101" divide:"100" json:"input_current"`              // [A] Charging equipment input current (Solar charge controller--PV array current)
	InputPower   float64 `addr:"0x3103" laddr:"0x3102" divide:"100" json:"input_power"` // [W] Charging equipment input power (Solar charge controller--PV array power)

	OutputVoltage float64 `addr:"0x3104" divide:"100" json:"output_voltage"`              // [V] Charging equipment output voltage (Battery charging voltage)
	OutputCurrent float64 `addr:"0x3105" divide:"100" json:"output_current"`              // [A] Charging equipment output current (Battery charging current)
	OutputPower   float64 `addr:"0x3107" laddr:"0x3106" divide:"100" json:"output_power"` // [W] Charging equipment output power (Battery charging power)

	LoadVoltage float64 `addr:"0x310C" divide:"100" json:"load_voltage"`              // [V] Disharging equipment output voltage (Load voltage)
	LoadCurrent float64 `addr:"0x310D" divide:"100" json:"load_current"`              // [A] Disharging equipment output current (Load current)
	LoadPower   float64 `addr:"0x310F" laddr:"0x310E" divide:"100" json:"load_power"` // [W] Disharging equipment output power (Load power)

	BatteryTemperature         float64 `addr:"0x3110" divide:"100" json:"battery_temperature"`          // [℃] Battery temperature in celtus
	EquipmentInsideTemperature float64 `addr:"0x3111" divide:"100" json:"equipment_inside_temperature"` // [℃] Temperature inside case
	PowerComponentsTemperature float64 `addr:"0x3112" divide:"100" json:"power_components_temperature"` // [℃] Heat sink surface temperature of equipment's power components

	BatterySOC              int     `addr:"0x311A" json:"battery_soc"`                            // [%] The percentage of battery's remaining capacity
	RemoteSensorTemperature float64 `addr:"0x311B" divide:"100" json:"remote_sensor_temperature"` // [℃] Remote temperature sensor (manual connect)
	CurrentSystemVoltage    float64 `addr:"0x311D" divide:"100" json:"current_system_voltage"`    // [V] Current system rated voltage. 1200, 2400, 3600, 4800 represent 12V，24V，36V，48V (Battery's real rated power)
}

type ControllerStatus struct {
	Battery                     string `addr:"0x3200" bits:"0000000000001111" enum:"00:Normal, 01:Overvolt, 02:Under volt, 03:Low volt disconnect, 04: Fault" json:"battery"`
	BatteryTemperature          string `addr:"0x3200" bits:"0000000011110000" enum:"00:Normal, 01:Over Temp (higher than the warning settings), 02:Low Temp (lower than the warning settings)" json:"battery_temperature"`
	BatteryInerternalResistance string `addr:"0x3200" bits:"0000000100000000" enum:"00:Normal, 01:Abnormanl" json:"battery_inerternal_resistance"`
	BatteryWrongVoltage         bool   `addr:"0x3200" bits:"1000000000000000" json:"battery_wrong_voltage"` // Wrong identification for rated voltage (i think this is about Current system rated voltage)

	Charging                           string `addr:"0x3201" bits:"0000000000000001" enum:"00:Standby, 01:Running" json:"charging"`
	ChargingStatus                     string `addr:"0x3201" bits:"0000000000000010" enum:"00:Normal, 01:Fault" json:"charging_status"`
	ChargingType                       string `addr:"0x3201" bits:"0000000000001100" enum:"00:No charging, 01:Float, 02:Boost, 03:Equalization" json:"charging_type"`
	ChargingPVInputIsShort             bool   `addr:"0x3201" bits:"0000000000010000" json:"charging_pv_input_is_short"`
	ChargingLoadMOSFETIsShort          bool   `addr:"0x3201" bits:"0000000010000000" json:"charging_load_mosfet_is_short"`
	ChargingTheLoadIsShort             bool   `addr:"0x3201" bits:"0000000100000000" json:"charging_the_load_is_short"`
	ChargingTheLoadIsOverCurrent       bool   `addr:"0x3201" bits:"0000001000000000" json:"charging_the_load_is_over_current"`
	ChargingInputIsOverCurrent         bool   `addr:"0x3201" bits:"0000010000000000" json:"charging_input_is_over_current"`
	ChargingAntiReverseMOSFETIsShort   bool   `addr:"0x3201" bits:"0000100000000000" json:"charging_anti_reverse_mosfet_is_short"`
	ChargingOrAntiReverseMOSFETIsShort bool   `addr:"0x3201" bits:"0001000000000000" json:"charging_or_anti_reverse_mosfet_is_short"`
	ChargingMOSFETIsShort              bool   `addr:"0x3201" bits:"0010000000000000" json:"charging_mosfet_is_short"`
	ChargingInputVoltStatus            string `addr:"0x3201" bits:"1100000000000000" enum:"00:Normal, 01:No power connected, 02:Higher volt input, 03:Input volt error" json:"charging_input_volt_status"`

	Discharging                            string `addr:"0x3202" bits:"0000000000000001" enum:"00:Standby, 01:Running" json:"discharging"`
	DischargingStatus                      string `addr:"0x3202" bits:"0000000000000010" enum:"00:Normal, 01:Fault" json:"discharging_status"`
	DischargingOutputOverpressure          string `addr:"0x3202" bits:"0000000000010000" enum:"00:No charging, 01:Float, 02:Boost, 03:Equalization" json:"discharging_output_overpressure"`
	DischargingBoostOverpressure           bool   `addr:"0x3202" bits:"0000000000100000" json:"discharging_boost_overpressure"`
	DischargingHighVoltageSideShortCircuit bool   `addr:"0x3202" bits:"0000000001000000" json:"discharging_high_voltage_side_short_circuit"`
	DischargingInputOverpressure           bool   `addr:"0x3202" bits:"0000000010000000" json:"discharging_input_overpressure"`
	DischargingOutputVoltageAbnormal       bool   `addr:"0x3202" bits:"0000000100000000" json:"discharging_output_voltage_abnormal"`
	DischargingUnableToStopDischarging     bool   `addr:"0x3202" bits:"0000001000000000" json:"discharging_unable_to_stop_discharging"`
	DischargingUnableToDischarge           bool   `addr:"0x3202" bits:"0000010000000000" json:"discharging_unable_to_discharge"`
	DischargingShortCircuit                bool   `addr:"0x3202" bits:"0000100000000000" json:"discharging_short_circuit"`
	DischargingOutputPower                 bool   `addr:"0x3202" bits:"0011000000000000" enum:"00:Light load, 01:Moderate, 02:Rated, 03:Overload" json:"discharging_output_power"`
	DischargingStat                        string `addr:"0x3202" bits:"1100000000000000" enum:"00:Normal, 01:Low, 02:High, 03:No access Input volt error" json:"discharging_stat"`
}

// Today - 00:00 Refresh every day
type ControllerStatistical struct {
	MaximumPVVoltageToday      float64 `addr:"0x3300" divide:"100" json:"maximum_pv_voltage_today"`
	MinimumPVVoltageToday      float64 `addr:"0x3301" divide:"100" json:"minimum_pv_voltage_today"`
	MaximumBatteryVoltageToday float64 `addr:"0x3302" divide:"100" json:"maximum_battery_voltage_today"`
	MinimumBatteryVoltageToday float64 `addr:"0x3303" divide:"100" json:"minimum_battery_voltage_today"`
	ConsumedEnergyToday        float64 `addr:"0x3305" laddr:"0x3304" divide:"100" json:"consumed_energy_today"`       // [KWH]
	ConsumedEnergyThisMonth    float64 `addr:"0x3307" laddr:"0x3306" divide:"100" json:"consumed_energy_this_month"`  // [KWH]
	ConsumedEnergyThisYear     float64 `addr:"0x3309" laddr:"0x3308" divide:"100" json:"consumed_energy_this_year"`   // [KWH]
	ConsumedEnergyTotal        float64 `addr:"0x330B" laddr:"0x330A" divide:"100" json:"consumed_energy_total"`       // [KWH]
	GeneratedEnergyToday       float64 `addr:"0x330D" laddr:"0x330C" divide:"100" json:"generated_energy_today"`      // [KWH]
	GeneratedEnergyThisMonth   float64 `addr:"0x330F" laddr:"0x330E" divide:"100" json:"generated_energy_this_month"` // [KWH]
	GeneratedEnergyThisYear    float64 `addr:"0x3311" laddr:"0x3310" divide:"100" json:"generated_energy_this_year"`  // [KWH]
	GeneratedEnergyTotal       float64 `addr:"0x3313" laddr:"0x3312" divide:"100" json:"generated_energy_total"`      // [KWH]
	BatteryVoltage             float64 `addr:"0x331A" divide:"100" json:"battery_voltage"`                            // [V]
	BatteryCurrent             float64 `addr:"0x331C" laddr:"0x331B" divide:"100" json:"battery_current"`             // [A]
}

type ControllerSettings struct {
	BatteryType            string  `addr:"0x9000" enum:"01:Sealed, 02:GEL, 03:Flooded, 00:User defined"`
	BatteryRatedCapacity   int     `addr:"0x9001"`              // [AH]
	TempCompensationCoeff  float64 `addr:"0x9002" divide:"100"` // Range 0-9 [mV/℃/2V]
	HighVoltageDisconnect  float64 `addr:"0x9003" divide:"100"` // [V]
	ChargingLimitVoltage   float64 `addr:"0x9004" divide:"100"` // [V]
	OverVoltageReconnect   float64 `addr:"0x9005" divide:"100"` // [V]
	EquilizationVoltage    float64 `addr:"0x9006" divide:"100"` // [V]
	BoostVoltageDisconnect float64 `addr:"0x9007" divide:"100"` // [V]
	FloatVoltageDisconnect float64 `addr:"0x9008" divide:"100"` // [V]
	BoostVoltageReconnect  float64 `addr:"0x9009" divide:"100"` // [V]

	LowVoltageReconnect     float64 `addr:"0x900A" divide:"100"` // [V]
	UnderVoltageRecover     float64 `addr:"0x900B" divide:"100"` // [V]
	UnderVoltageWarning     float64 `addr:"0x900C" divide:"100"` // [V]
	LowVoltageDisconnect    float64 `addr:"0x900D" divide:"100"` // [V]
	DischargingLimitVoltage float64 `addr:"0x900E" divide:"100"` // [V]
	RealTimeClock1          int     `addr:"0x9013"`              // D7-0 Sec, D15-8 Min.(Year, Month, Day, Min, Sec. should be written simultaneously)
	RealTimeClock2          int     `addr:"0x9014"`              // D7-0 Hour, D15-8 Day
	RealTimeClock3          int     `addr:"0x9015"`              //D7-0 Month, D15-8 Year

	BatteryTemperatureWarningUpperLimit         float64 `addr:"0x9017" divide:"100"` // [℃]
	BatteryTemperatureWarningLowerLimit         float64 `addr:"0x9018" divide:"100"` // [℃]
	ControllerInnerTemperatureUpperLimit        float64 `addr:"0x9019" divide:"100"` // [℃]
	ControllerInnerTemperatureUpperLimitRecover float64 `addr:"0x901A" divide:"100"` // [℃]

	DayTimeThresholdVolt    float64 `addr:"0x901E" divide:"100"` // (DTTV) PV lower than this value, controller would detect it as sundown [V]
	LightSignalStartupDelay int     `addr:"0x901F"`              // PV voltage lower than NTTV, and duration exceeds the Light signal startup (night) delay time, controller would detect it as night time. [min]
	LightTimeThresholdVolt  float64 `addr:"0x9020" divide:"100"` // (NTTV) PV voltage higher than this value, controller would detect it as sunrise [V]
	LightSignalCloseDelay   int     `addr:"0x9021"`              // PV voltage higher than DTTV, and duration exceeds the Light signal close (day) delay time, controller would detect it as day time. [min]

	LoadControllingMode string `addr:"0x903D" enum:"00:Manual Control, 01:Light ON/OFF, 02:Light ON+ Timer, 03:Time Control"`
	WorkingTimeLength1  int    `addr:"0x903E"` //The length of load output timer1, D15-D8,hour, D7-D0, minute
	WorkingTimeLength2  int    `addr:"0x903F"` //The length of load output timer2, D15-D8,hour, D7-D0, minute

	TurnOnTiming1Sec  int `addr:"0x9042"` // Turn on/off timing of load output. [sec]
	TurnOnTiming1Min  int `addr:"0x9043"` // Turn on/off timing of load output. [min]
	TurnOnTiming1Hour int `addr:"0x9044"` // Turn on/off timing of load output. [hour]

	TurnOffTiming1Sec  int `addr:"0x9045"` // Turn on/off timing of load output. [sec]
	TurnOffTiming1Min  int `addr:"0x9046"` // Turn on/off timing of load output. [min]
	TurnOffTiming1Hour int `addr:"0x9047"` // Turn on/off timing of load output. [hour]

	TurnOnTiming2Sec  int `addr:"0x9048"` // Turn on/off timing of load output. [sec]
	TurnOnTiming2Min  int `addr:"0x9049"` // Turn on/off timing of load output. [min]
	TurnOnTiming2Hour int `addr:"0x904A"` // Turn on/off timing of load output. [hour]

	TurnOffTiming2Sec  int `addr:"0x904B"` // Turn on/off timing of load output. [sec]
	TurnOffTiming2Min  int `addr:"0x904C"` // Turn on/off timing of load output. [min]
	TurnOffTiming2Hour int `addr:"0x904D"` // Turn on/off timing of load output. [hour]

	BacklightTime int `addr:"0x9063"` // Close after LCD backlight light setting the number of secends [S]
	LengthOfNigh  int `addr:"0x9065"` // Set default values of the whole night length of time. D15-D8,hour, D7-D0, minute [S]
	// DeviceConfigureOfMainPowerSupply string `addr:"0x9066" enum:"01:Battery is main, 02:AC-DC power mainly"`
	BatteryRatedVoltageCode      string `addr:"0x9067" enum:"0:Auto recognize, 1:12V, 2:24V, 3:36V, 4:48V, 5:60V, 6:110V, 7:120V, 8: 220V, 9: 240V"`
	DefaultLoadOnOffInManualMode string `addr:"0x906A" enum:"0:Off, 1:On"`
	EqualizeDuration             int    `addr:"0x906B"` // Usually 0-120 minutes [min]
	BoostDuration                int    `addr:"0x906C"` // Usually 0-120 minutes [min]
	DischargingPercentage        int    `addr:"0x906D"` // Usually 20%-80%. The percentage of battery's remaining capacity when stop charging [%]
	ChargingPercentage           int    `addr:"0x906E"` // Depth of charge, 100% [%]

	ManagementModesOfBatteryChargingAndDischarging string `addr:"0x9070" enum:"0:Voltage compensation, 1:SOC"` // Management modes of battery charge and discharge, voltage compensation : 0 and SOC : 1
}

type ControllerSwitches struct {
	ChargingDevice            bool `addr:"0x0" mode:"raw"`  // 1 Charging device on 0 Charging device off
	OutputControlMode         bool `addr:"0x1" mode:"raw"`  // 1 Output control mode manual 0 Output control mode automatic
	ManualControlTheLoad      bool `addr:"0x2" mode:"raw"`  // When the load is manual mode，1-manual on 0 -manual off
	DefaultControlTheLoad     bool `addr:"0x3" mode:"raw"`  // When the load is default mode，1-manual on 0 -manual off
	EnableLoadTestMode        bool `addr:"0x5" mode:"raw"`  // 1 Enable 0 Disable(normal)
	ForceTheLoad              bool `addr:"0x6" mode:"raw"`  // 1 Turn on 0 Turn off (used for temporary test of the load）
	RestoresystemDefaults     bool `addr:"0x13" mode:"raw"` // 1 yes 0 no
	ClearGeneratingStatistics bool `addr:"0x14" mode:"raw"` // 1 clear. Root privileges to perform
}

type ControllerDiscrete struct {
	OverTemperatureInsideTheDevice string `addr:"0x2000" mode:"raw" enum:"0:Normal, 1:Higher than the over-temperature protection point"` // 1 The temperature inside the controller is higher than the over-temperature protection point. 0 Normal
	DayNight                       string `addr:"0x200C" mode:"raw" enum:"0:Day, 1:Night"`                                                // 1-Night, 0-Day
}

type ControllerSummary struct {
	ControllerRated       *ControllerRated       `json:"controller_rated,omitempty"`
	ControllerRealTime    *ControllerRealTime    `json:"controller_real_time,omitempty"`
	ControllerStatus      *ControllerStatus      `json:"controller_status,omitempty"`
	ControllerStatistical *ControllerStatistical `json:"controller_statistical,omitempty"`
	ControllerSettings    *ControllerSettings    `json:"controller_settings,omitempty"`
	ControllerSwitches    *ControllerSwitches    `json:"controller_switches,omitempty"`
	ControllerDiscrete    *ControllerDiscrete    `json:"controller_discrete,omitempty"`
}
