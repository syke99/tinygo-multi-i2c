package multi

import (
	"time"

	"tinygo.org/x/drivers"
)

type Bh1750 struct {
	bus     drivers.I2C
	Address uint16
	mode    byte
}

func newBh1750(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Bh1750{
			bus:     bus,
			Address: addr,
			mode:    BH1750_CONTINUOUS_HIGH_RES_MODE,
		}
	} else {
		return Bh1750{
			bus:     bus,
			Address: Bh1750Address,
			mode:    BH1750_CONTINUOUS_HIGH_RES_MODE,
		}
	}
}

func (d Bh1750) configure() error {
	if err := d.bus.Tx(d.Address, []byte{BH1750_POWER_ON}, nil); err != nil {
		return err
	}
	d.SetMode(d.mode)

	return nil
}

func (d Bh1750) RawSensorData() uint16 {

	buf := []byte{1, 0}
	d.bus.Tx(d.Address, nil, buf)
	return (uint16(buf[0]) << 8) | uint16(buf[1])
}

func (d Bh1750) Illuminance() int32 {

	lux := uint32(d.RawSensorData())
	var coef uint32
	if d.mode == BH1750_CONTINUOUS_HIGH_RES_MODE || d.mode == BH1750_ONE_TIME_HIGH_RES_MODE {
		coef = BH1750_HIGH_RES
	} else if d.mode == BH1750_CONTINUOUS_HIGH_RES_MODE_2 || d.mode == BH1750_ONE_TIME_HIGH_RES_MODE_2 {
		coef = BH1750_HIGH_RES2
	} else {
		coef = BH1750_LOW_RES
	}
	// 100 * coef * lux * (5/6)
	// 5/6 = measurement accuracy as per the datasheet
	return int32(250 * coef * lux / 3)
}

func (d Bh1750) SetMode(mode byte) {
	d.mode = mode
	d.bus.Tx(d.Address, []byte{byte(d.mode)}, nil)
	time.Sleep(10 * time.Millisecond)
}
