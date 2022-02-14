package multi

import "tinygo.org/x/drivers"

//-------------------------------------------------------------------------------------
// Lps22hb
type Lps22hb struct {
	bus     drivers.I2C
	Address uint8
}

//-------------------------------------------------------------------------------------
// Lps22hb
func newLps22hb(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Lps22hb{
			bus:     bus,
			Address: uint8(addr),
		}
	} else {
		return Lps22hb{
			bus:     bus,
			Address: Lps22hbAddress,
		}
	}
}

//-------------------------------------------------------------------------------------
// Lps22hb
func (d Lps22hb) configure() {
	// Lps22hb doesn't need a configure method??
}

//-------------------------------------------------------------------------------------
// Lps22hb
func (d Lps22hb) connected() bool {
	data := []byte{0}
	d.bus.ReadRegister(d.Address, LPS22HB_WHO_AM_I_REG, data)
	return data[0] == 0xB1
}

// ReadPressure returns the pressure in milli pascals (mPa).
func (d Lps22hb) ReadPressure() (pressure int32, err error) {
	d.waitForOneShot()

	// read data
	data := []byte{0, 0, 0}
	d.bus.ReadRegister(d.Address, LPS22HB_PRESS_OUT_REG, data[:1])
	d.bus.ReadRegister(d.Address, LPS22HB_PRESS_OUT_REG+1, data[1:2])
	d.bus.ReadRegister(d.Address, LPS22HB_PRESS_OUT_REG+2, data[2:])
	pValue := float32(uint32(data[2])<<16|uint32(data[1])<<8|uint32(data[0])) / 4096.0

	return int32(pValue * 1000), nil
}

// ReadTemperature returns the temperature in celsius milli degrees (°C/1000).
func (d Lps22hb) ReadTemperature() (temperature int32, err error) {
	d.waitForOneShot()

	// read data
	data := []byte{0, 0}
	d.bus.ReadRegister(d.Address, LPS22HB_TEMP_OUT_REG, data[:1])
	d.bus.ReadRegister(d.Address, LPS22HB_TEMP_OUT_REG+1, data[1:])
	tValue := float32(int16(uint16(data[1])<<8|uint16(data[0]))) / 100.0

	return int32(tValue * 1000), nil
}

// private functions

// wait and trigger one shot in block update
func (d Lps22hb) waitForOneShot() {
	// trigger one shot
	d.bus.WriteRegister(d.Address, LPS22HB_CTRL2_REG, []byte{0x01})

	// wait until one shot is cleared
	data := []byte{1}
	for {
		d.bus.ReadRegister(d.Address, LPS22HB_CTRL2_REG, data)
		if data[0]&0x01 == 0 {
			break
		}
	}
}
