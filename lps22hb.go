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
