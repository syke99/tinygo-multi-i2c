package multi

import "tinygo.org/x/drivers"

//-------------------------------------------------------------------------------------
// Sht3x
type Sht3x struct {
	bus     drivers.I2C
	Address uint16
}

//-------------------------------------------------------------------------------------
// Sht3x
func newSht3x(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Sht3x{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Sht3x{
			bus:     bus,
			Address: Sht3xAddressA,
		}
	}
}

//-------------------------------------------------------------------------------------
// Sht3x
func (d Sht3x) configure() {
	// Sht3x doesn't need a configure method??
}
