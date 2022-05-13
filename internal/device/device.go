package device

import (
	"github.com/goburrow/modbus"
	log "github.com/sirupsen/logrus"
)

type Device struct {
	rtu    *modbus.RTUClientHandler
	client modbus.Client
}

func (d *Device) Connect() error {
	err := d.rtu.Connect()
	if err != nil {
		log.Errorln("RTU Connect: %v", err)
		return err
	}

	d.client = modbus.NewClient(d.rtu)
	return nil
}

func (d *Device) Close() {
	d.rtu.Close()
}

func (d *Device) IsConnected() bool {
	return d.client != nil && d.rtu.Connect() == nil
}
