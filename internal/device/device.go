package device

import (
	"github.com/goburrow/modbus"
)

type Device struct {
	rtu    *modbus.RTUClientHandler
	client modbus.Client
}

func (d *Device) Connect() error {
	err := d.rtu.Connect()
	if err != nil {
		return err
	}

	d.client = modbus.NewClient(d.rtu)
	return nil
}

func (d *Device) Close() {
	d.client = nil
	d.rtu.Close()
}

func (d *Device) IsConnected() bool {
	return d.client != nil && d.rtu.Connect() == nil
}
