package multi

import "tinygo.org/x/drivers"

//-------------------------------------------------------------------------------------
// Lsm6ds3
type lsm6ds3AccelRange uint8
type lsm6ds3AccelSampleRate uint8
type lsm6ds3AccelBandwidth uint8

type lsm6ds3GyroRange uint8
type lsm6ds3GyroSampleRate uint8

type Lsm6ds3 struct {
	bus             drivers.I2C
	Address         uint16
	accelRange      lsm6ds3AccelRange
	accelSampleRate lsm6ds3AccelSampleRate
	accelBandWidth  lsm6ds3AccelBandwidth
	gyroRange       lsm6ds3GyroRange
	gyroSampleRate  lsm6ds3GyroSampleRate
	dataBufferSix   []uint8
	dataBufferTwo   []uint8
}

//-------------------------------------------------------------------------------------
// Lsm6ds3
func newLsm6ds3(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Lsm6ds3{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Lsm6ds3{
			bus:     bus,
			Address: Lsm6ds3Address,
		}
	}
}
