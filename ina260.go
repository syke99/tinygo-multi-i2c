package multi

import "tinygo.org/x/drivers"

//-------------------------------------------------------------------------------------
// Ina260
type Ina260 struct {
	bus     drivers.I2C
	Address uint16
}

//-------------------------------------------------------------------------------------
// Ina260
func newIna260(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Ina260{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Ina260{
			bus:     bus,
			Address: Ina250Address,
		}
	}
}

//-------------------------------------------------------------------------------------
// Ina260
func (d Ina260) configure() {
	// Ina260 may need to be removed?
}

//-------------------------------------------------------------------------------------
// Ina260
func (i Ina260) connected() error {
	// Ina260 may need to be removed?
	return nil
}
