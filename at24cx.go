package multi

import "tinygo.org/x/drivers"

//-------------------------------------------------------------------------------------
// AT24CX
type At24cx struct {
	bus               drivers.I2C
	Address           uint16
	pageSize          uint16
	currentRAMAddress uint16
	startRAMAddress   uint16
	endRAMAddress     uint16
}

//-------------------------------------------------------------------------------------
// AT24CX
func newAt24cx(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return At24cx{
			bus:     bus,
			Address: addr,
		}
	} else {
		return At24cx{
			bus:     bus,
			Address: At24cxAddress,
		}
	}
}

//-------------------------------------------------------------------------------------
// AT24CX
func (d At24cx) configure() {
	// At24cx may need to be removed?
}

//-------------------------------------------------------------------------------------
// AT24CX
// skip for now
