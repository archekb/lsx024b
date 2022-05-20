package device

import (
	"time"

	"github.com/archekb/lsx024b/internal/models"
	"github.com/goburrow/modbus"
)

type DeviceLSB struct {
	Device
}

func NewLSB(port string, id int) *DeviceLSB {
	rtu := modbus.NewRTUClientHandler(port)
	rtu.BaudRate = 115200
	rtu.DataBits = 8
	rtu.Parity = "N"
	rtu.StopBits = 1
	rtu.Timeout = 5 * time.Second

	if id >= 1 && id <= 247 {
		rtu.SlaveId = byte(id)
	} else {
		rtu.SlaveId = 1
	}

	return &DeviceLSB{Device: Device{rtu: rtu}}
}

func (d *DeviceLSB) GetRated() (*models.LSB_Rated, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSB_Rated{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSB) GetRealTime() (*models.LSB_RealTime, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSB_RealTime{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		return nil, err
	}
	return &data, nil
}

func (d *DeviceLSB) GetStatus() (*models.LSB_Status, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSB_Status{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		return nil, err
	}
	return &data, nil
}

func (d *DeviceLSB) GetStatistical() (*models.LSB_Statistical, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSB_Statistical{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSB) GetSettings() (*models.LSB_Settings, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSB_Settings{}
	if err := FillStruct(&data, d.client.ReadHoldingRegisters); err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSB) GetSwitches() (*models.LSB_Switches, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSB_Switches{}
	if err := FillStruct(&data, d.client.ReadCoils); err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSB) GetDiscrete() (*models.LSB_Discrete, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSB_Discrete{}
	if err := FillStruct(&data, d.client.ReadDiscreteInputs); err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSB) Summary() (*models.SummaryLSB, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.SummaryLSB{}
	var err error

	data.Rated, err = d.GetRated()
	if err != nil {
		return nil, err
	}

	data.RealTime, err = d.GetRealTime()
	if err != nil {
		return nil, err
	}

	data.Status, err = d.GetStatus()
	if err != nil {
		return nil, err
	}

	data.Statistical, err = d.GetStatistical()
	if err != nil {
		return nil, err
	}

	data.Settings, err = d.GetSettings()
	if err != nil {
		return nil, err
	}

	// data.Switches, err = d.GetSwitches()
	// if err != nil {
	// 	return nil, err
	// }

	// data.Discrete, err = d.GetDiscrete()
	// if err != nil {
	// 	return nil, err
	// }

	return &data, nil
}

// func (d *Device) Test() {
// 	//
// 	dataDiscrete := models.Discrete{}
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
