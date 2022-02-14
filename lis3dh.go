package multi

type Lis3dhRange uint8
type Lis3dhRate uint8

type Lis3dh struct {
	bus     I2C
	Address uint16
	r       Lis3dhRange
}

func newLis3dh(bus I2C, addr uint16) interface{} {
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

func (d Lis3dh) connected() bool {
	data := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), LIS3DH_WHO_AM_I, data)
	if err != nil {
		return false
	}
	return data[0] == 0x33
}

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

// SetRange sets the G range for LIS3DH.
func (d Lis3dh) SetRange(r Lis3dhRange) {
	ctl := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), LIS3DH_REG_CTRL4, ctl)
	if err != nil {
		println(err.Error())
	}
	// mask off bits
	ctl[0] &^= 0x30
	ctl[0] |= (byte(r) << 4)
	d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL4, ctl)

	// store the new range
	d.r = r
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

// ReadAcceleration reads the current acceleration from the device and returns
// it in µg (micro-gravity). When one of the axes is pointing straight to Earth
// and the sensor is not moving the returned value will be around 1000000 or
// -1000000.
func (d Lis3dh) ReadAcceleration() (int32, int32, int32, error) {
	x, y, z := d.ReadRawAcceleration()
	divider := float32(1)
	switch d.r {
	case LIS3DH_RANGE_16_G:
		divider = 1365
	case LIS3DH_RANGE_8_G:
		divider = 4096
	case LIS3DH_RANGE_4_G:
		divider = 8190
	case LIS3DH_RANGE_2_G:
		divider = 16380
	}

	return int32(float32(x) / divider * 1000000), int32(float32(y) / divider * 1000000), int32(float32(z) / divider * 1000000), nil
}

// ReadRawAcceleration returns the raw x, y and z axis from the LIS3DH
func (d Lis3dh) ReadRawAcceleration() (x int16, y int16, z int16) {
	d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_OUT_X_L|0x80, nil)

	data := []byte{0, 0, 0, 0, 0, 0}
	d.bus.Tx(d.Address, nil, data)

	x = int16((uint16(data[1]) << 8) | uint16(data[0]))
	y = int16((uint16(data[3]) << 8) | uint16(data[2]))
	z = int16((uint16(data[5]) << 8) | uint16(data[4]))

	return
}
