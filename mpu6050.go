package multi

import "tinygo.org/x/drivers"

//-------------------------------------------------------------------------------------
// Mpu6050
type Mpu6050 struct {
	bus     drivers.I2C
	Address uint16
}

//-------------------------------------------------------------------------------------
// Mpu6050
func newMpu6050(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Mpu6050{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Mpu6050{
			bus:     bus,
			Address: Mpu6050Address,
		}
	}
}

//-------------------------------------------------------------------------------------
// Mpu6050
func (d Mpu6050) configure() {
	d.bus.WriteRegister(uint8(d.Address), MPU6050_PWR_MGMT_1, []uint8{0})
}

//-------------------------------------------------------------------------------------
// Mpu6050
func (d Mpu6050) connected() bool {
	data := []byte{0}
	d.bus.ReadRegister(uint8(d.Address), MPU6050_WHO_AM_I, data)
	return data[0] == 0x68
}
