package multi

import "tinygo.org/x/drivers"

//-------------------------------------------------------------------------------------
// Lis3dh
type Lis3dh struct {
	bus     drivers.I2C
	Address uint16
	r       Lis3dhRange
}

//-------------------------------------------------------------------------------------
// Lis3dh
func newLis3dh(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Lis3dh{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Lis3dh{
			bus:     bus,
			Address: Lis3dhAddress0,
		}
	}
}

//-------------------------------------------------------------------------------------
// Lis3dh
func (d Lis3dh) configure() error {
	// enable all axes, normal mode
	if err := d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL1, []byte{0x07}); err != nil {
		return err
	}

	// 400Hz rate
	d.SetDataRate(LIS3DH_DATARATE_400_HZ)

	// High res & BDU enabled
	if err := d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL4, []byte{0x88}); err != nil {
		return err
	}

	// get current range
	d.r = d.ReadRange()

	return nil
}

//-------------------------------------------------------------------------------------
// Lis3dh
func (d Lis3dh) connected() bool {
	data := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), LIS3DH_WHO_AM_I, data)
	if err != nil {
		return false
	}
	return data[0] == 0x33
}

//-------------------------------------------------------------------------------------
// Lis3dh
func (d Lis3dh) SetDataRate(rate Lis3dhDataRate) {
	ctl1 := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), LIS3DH_REG_CTRL1, ctl1)
	if err != nil {
		println(err.Error())
	}
	// mask off bits
	ctl1[0] &^= 0xf0
	ctl1[0] |= (byte(rate) << 4)
	d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL1, ctl1)
}

func (d Lis3dh) ReadRange() (r Lis3dhRange) {
	ctl := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), LIS3DH_REG_CTRL4, ctl)
	if err != nil {
		println(err.Error())
	}
	// mask off bits
	r = Lis3dhRange(ctl[0] >> 4)
	r &= 0x03

	return r
}
