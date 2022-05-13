package dev

import (
	"time"

	"github.com/goburrow/modbus"
	log "github.com/sirupsen/logrus"
)

type Dev struct {
	rtu    *modbus.RTUClientHandler
	client modbus.Client
}

func New(port string) (*Dev, error) {
	rtu := modbus.NewRTUClientHandler(port)
	rtu.SlaveId = 1
	rtu.BaudRate = 115200
	rtu.DataBits = 8
	rtu.Parity = "N"
	rtu.StopBits = 1
	rtu.Timeout = 5 * time.Second

	err := rtu.Connect()
	if err != nil {
		log.Errorln("RTU Connect: %v", err)
	}

	client := modbus.NewClient(rtu)
	return &Dev{
		rtu:    rtu,
		client: client,
	}, nil
}

func (d *Dev) Close() {
	d.rtu.Close()
}

func (d *Dev) ControllerRated() (*ControllerRated, error) {
	data := ControllerRated{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *Dev) ControllerRealTime() (*ControllerRealTime, error) {
	data := ControllerRealTime{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *Dev) ControllerStatus() (*ControllerStatus, error) {
	data := ControllerStatus{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *Dev) ControllerStatistical() (*ControllerStatistical, error) {
	data := ControllerStatistical{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *Dev) ControllerSettings() (*ControllerSettings, error) {
	data := ControllerSettings{}
	if err := FillStruct(&data, d.client.ReadHoldingRegisters); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *Dev) ControllerSwitches() (*ControllerSwitches, error) {
	data := ControllerSwitches{}
	if err := FillStruct(&data, d.client.ReadCoils); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *Dev) ControllerDiscrete() (*ControllerDiscrete, error) {
	data := ControllerDiscrete{}
	if err := FillStruct(&data, d.client.ReadDiscreteInputs); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *Dev) ControllerSummary() (*ControllerSummary, error) {
	data := ControllerSummary{}
	var err error

	data.ControllerRated, err = d.ControllerRated()
	if err != nil {
		return nil, err
	}

	data.ControllerRealTime, err = d.ControllerRealTime()
	if err != nil {
		return nil, err
	}

	data.ControllerStatus, err = d.ControllerStatus()
	if err != nil {
		return nil, err
	}

	data.ControllerStatistical, err = d.ControllerStatistical()
	if err != nil {
		return nil, err
	}

	data.ControllerSettings, err = d.ControllerSettings()
	if err != nil {
		return nil, err
	}

	// data.ControllerSwitches, err = d.ControllerSwitches()
	// if err != nil {
	// 	return nil, err
	// }

	// data.ControllerDiscrete, err = d.ControllerDiscrete()
	// if err != nil {
	// 	return nil, err
	// }

	return &data, nil
}

// func (d *Dev) Test() {
// 	//
// 	dataDiscrete := ControllerDiscrete{}
// 	err = FillStruct(&dataDiscrete, client.ReadDiscreteInputs)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Printf("%+v\n\n", dataDiscrete)
// 	var ChargingState1 int = 1
// 	log.Println(ChargingState1&1, (ChargingState1&2)>>1, (ChargingState1&12)>>2, (ChargingState1&16)>>4, (ChargingState1&128)>>7, (ChargingState1&256)>>8, (ChargingState1&1024)>>10, (ChargingState1&2048)>>11, (ChargingState1&4096)>>12, (ChargingState1&8192)>>13, (ChargingState1&49152)>>14)

// 	// //client.WriteSingleCoil()
// 	res, err := client.WriteSingleCoil(0x2, 0x0000) // 0x0000)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(res, Decode(res), string(res))

// 	// res, err := client.WriteMultipleRegisters(0x903D, 0x01, Encode([]uint16{0}))
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }
// 	// log.Println(res, Decode(res), string(res))

// }
