package multi

import "tinygo.org/x/drivers"

//-------------------------------------------------------------------------------------
// Ds3231
type Ds3231Mode uint8

type Ds3231 struct {
	bus     drivers.I2C
	Address uint16
}

//-------------------------------------------------------------------------------------
// Ds3231
func newDs3231(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Ds3231{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Ds3231{
			bus:     bus,
			Address: Ds3231Address,
		}
	}
}

//-------------------------------------------------------------------------------------
// Ds3231
func (d Ds3231) configure() {
	// DS3231 doesn't need a configure method??
}
