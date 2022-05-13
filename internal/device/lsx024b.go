package device

import (
	"time"

	"github.com/archekb/lsx024b/internal/models"
	"github.com/goburrow/modbus"
	log "github.com/sirupsen/logrus"
)

type DeviceLSX024B struct {
	Device
}

func NewLSX024B(port string, id int) *DeviceLSX024B {
	rtu := modbus.NewRTUClientHandler(port)
	rtu.SlaveId = byte(id)
	rtu.BaudRate = 115200
	rtu.DataBits = 8
	rtu.Parity = "N"
	rtu.StopBits = 1
	rtu.Timeout = 5 * time.Second

	if id > 0 && id < 256 {
		rtu.SlaveId = byte(id)
	} else {
		rtu.SlaveId = 1
	}

	return &DeviceLSX024B{Device: Device{rtu: rtu}}
}

func (d *DeviceLSX024B) GetRated() (*models.LSX024BRated, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSX024BRated{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		log.Println(err)
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSX024B) GetRealTime() (*models.LSX024BRealTime, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSX024BRealTime{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *DeviceLSX024B) GetStatus() (*models.LSX024BStatus, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSX024BStatus{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *DeviceLSX024B) GetStatistical() (*models.LSX024BStatistical, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSX024BStatistical{}
	if err := FillStruct(&data, d.client.ReadInputRegisters); err != nil {
		log.Println(err)
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSX024B) GetSettings() (*models.LSX024BSettings, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSX024BSettings{}
	if err := FillStruct(&data, d.client.ReadHoldingRegisters); err != nil {
		log.Println(err)
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSX024B) GetSwitches() (*models.LSX024BSwitches, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSX024BSwitches{}
	if err := FillStruct(&data, d.client.ReadCoils); err != nil {
		log.Println(err)
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSX024B) GetDiscrete() (*models.LSX024BDiscrete, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.LSX024BDiscrete{}
	if err := FillStruct(&data, d.client.ReadDiscreteInputs); err != nil {
		log.Println(err)
		return nil, err
	}

	return &data, nil
}

func (d *DeviceLSX024B) Summary() (*models.SummaryLSX024B, error) {
	if !d.IsConnected() {
		return nil, ErrNotConnected
	}

	data := models.SummaryLSX024B{}
	var err error

	data.LSX024BRated, err = d.GetRated()
	if err != nil {
		return nil, err
	}

	data.LSX024BRealTime, err = d.GetRealTime()
	if err != nil {
		return nil, err
	}

	data.LSX024BStatus, err = d.GetStatus()
	if err != nil {
		return nil, err
	}

	data.LSX024BStatistical, err = d.GetStatistical()
	if err != nil {
		return nil, err
	}

	data.LSX024BSettings, err = d.GetSettings()
	if err != nil {
		return nil, err
	}

	// data.LSX024BSwitches, err = d.GetSwitches()
	// if err != nil {
	// 	return nil, err
	// }

	// data.LSX024BDiscrete, err = d.GetDiscrete()
	// if err != nil {
	// 	return nil, err
	// }

	return &data, nil
}

// func (d *Device) Test() {
// 	//
// 	dataDiscrete := models.LSX024BDiscrete{}
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
