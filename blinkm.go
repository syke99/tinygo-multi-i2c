package multi

import "tinygo.org/x/drivers"

type Blinkm struct {
	bus     drivers.I2C
	Address uint16
}

func newBlinkm(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Blinkm{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Blinkm{
			bus:     bus,
			Address: BlinkmAddress,
		}
	}
}

func (d Blinkm) configure() error {
	if err := d.bus.Tx(d.Address, []byte{'o'}, nil); err != nil {
		return err
	}

	return nil
}

func (d Blinkm) Version() (major, minor byte, err error) {
	version := []byte{0, 0}
	d.bus.Tx(d.Address, []byte{BLINKM_GET_FIRMWARE}, version)
	return version[0], version[1], nil
}

func (d Blinkm) SetRGB(r, g, b byte) error {
	d.bus.Tx(d.Address, []byte{BLINKM_TO_RGB, r, g, b}, nil)
	return nil
}

func (d Blinkm) GetRGB() (r, g, b byte, err error) {
	color := []byte{0, 0, 0}
	d.bus.Tx(d.Address, []byte{BLINKM_GET_RGB}, color)
	return color[0], color[1], color[2], nil
}

func (d Blinkm) FadeToRGB(r, g, b byte) error {
	d.bus.Tx(d.Address, []byte{BLINKM_FADE_TO_RGB, r, g, b}, nil)
	return nil
}

func (d Blinkm) StopScript() error {
	d.bus.Tx(d.Address, []byte{BLINKM_STOP_SCRIPT}, nil)
	return nil
}
