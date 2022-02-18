package v2

type Mpu6050 struct {
	bus     I2C
	Address uint16
}

func newMpu6050(bus I2C, addr uint16) Mpu6050 {
	if addr != 0 {
		return Mpu6050{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Mpu6050{
			bus:     bus,
			Address: Mpu6050Address_1,
		}
	}
}

func (d Mpu6050) configure() {
	d.bus.WriteRegister(uint8(d.Address), MPU6050_PWR_MGMT_1, []uint8{0})
}

func (d Mpu6050) connected() bool {
	data := []byte{0}
	d.bus.ReadRegister(uint8(d.Address), MPU6050_WHO_AM_I, data)
	return data[0] == 0x68
}

// ReadAcceleration reads the current acceleration from the device and returns
// it in µg (micro-gravity). When one of the axes is pointing straight to Earth
// and the sensor is not moving the returned value will be around 1000000 or
// -1000000.
func (d Mpu6050) ReadAcceleration() (x int32, y int32, z int32) {
	data := make([]byte, 6)
	d.bus.ReadRegister(uint8(d.Address), MPU6050_ACCEL_XOUT_H, data)
	// Now do two things:
	// 1. merge the two values to a 16-bit number (and cast to a 32-bit integer)
	// 2. scale the value to bring it in the -1000000..1000000 range.
	//    This is done with a trick. What we do here is essentially multiply by
	//    1000000 and divide by 16384 to get the original scale, but to avoid
	//    overflow we do it at 1/64 of the value:
	//      1000000 / 64 = 15625
	//      16384   / 64 = 256
	x = int32(int16((uint16(data[0])<<8)|uint16(data[1]))) * 15625 / 256
	y = int32(int16((uint16(data[2])<<8)|uint16(data[3]))) * 15625 / 256
	z = int32(int16((uint16(data[4])<<8)|uint16(data[5]))) * 15625 / 256
	return
}

// ReadRotation reads the current rotation from the device and returns it in
// µ°/s (micro-degrees/sec). This means that if you were to do a complete
// rotation along one axis and while doing so integrate all values over time,
// you would get a value close to 360000000.
func (d Mpu6050) ReadRotation() (x int32, y int32, z int32) {
	data := make([]byte, 6)
	d.bus.ReadRegister(uint8(d.Address), MPU6050_GYRO_XOUT_H, data)
	// First the value is converted from a pair of bytes to a signed 16-bit
	// value and then to a signed 32-bit value to avoid integer overflow.
	// Then the value is scaled to µ°/s (micro-degrees per second).
	// This is done in the following steps:
	// 1. Multiply by 250 * 1000_000
	// 2. Divide by 32768
	// The following calculation (x * 15625 / 2048 * 1000) is essentially the
	// same but avoids overflow. First both operations are divided by 16 leading
	// to multiply by 15625000 and divide by 2048, and then part of the multiply
	// is done after the divide instead of before.
	x = int32(int16((uint16(data[0])<<8)|uint16(data[1]))) * 15625 / 2048 * 1000
	y = int32(int16((uint16(data[2])<<8)|uint16(data[3]))) * 15625 / 2048 * 1000
	z = int32(int16((uint16(data[4])<<8)|uint16(data[5]))) * 15625 / 2048 * 1000
	return
}
