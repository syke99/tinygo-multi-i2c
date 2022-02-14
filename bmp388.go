package multi

import (
	"errors"

	"tinygo.org/x/drivers"
)

//-------------------------------------------------------------------------------------
// // BMP388
type bmp388CalibrationCoefficients struct {
	// Temperature compensation
	t1 uint16
	t2 uint16
	t3 int8

	// Pressure compensation
	p1  int16
	p2  int16
	p3  int8
	p4  int8
	p5  uint16
	p6  uint16
	p7  int8
	p8  int8
	p9  int16
	p10 int8
	p11 int8
}

type bmp388Oversampling byte
type bmp388Mode byte
type bmp388OutputDataRate byte
type bmp388FilterCoefficient byte

type bmp388Config struct {
	Pressure    byte
	Temperature bmp388Oversampling
	Mode        bmp388Mode
	ODR         bmp388OutputDataRate
	IIR         bmp388FilterCoefficient
}

type Bmp388 struct {
	bus     drivers.I2C
	Address uint8
	cali    bmp388CalibrationCoefficients
	Config  bmp388Config
}

// These variables belong to BMP388
var (
	errConfigWrite  = errors.New("bmp388: failed to configure sensor, check connection")
	errConfig       = errors.New("bmp388: there is a problem with the configuration, try reducing ODR")
	errCaliRead     = errors.New("bmp388: failed to read calibration coefficient register")
	errSoftReset    = errors.New("bmp388: failed to perform a soft reset")
	errNotConnected = errors.New("bmp388: not connected")
)

//-------------------------------------------------------------------------------------
// // BMP388
func newBmp388(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Bmp388{
			bus:     bus,
			Address: uint8(addr),
		}
	} else {
		return Bmp388{
			bus:     bus,
			Address: Bmp388Address,
		}
	}
}

//-------------------------------------------------------------------------------------
// // BMP388
func (d Bmp388) configure() {
	// Bmp388 may need to be removed?
}

//-------------------------------------------------------------------------------------
// // BMP388
// func (d Bmp388) connected() bool {
// 	data, err := d.readRegister(BMP388_RegChipId, 1)
// 	return err == nil && data[0] == BMP388_ChipId // returns true if i2c comm was good and response equals 0x50
// }

//-------------------------------------------------------------------------------------
// // BMP388
// func (d Bmp388) tlinCompensate() (int64, error) {
// 	rawTemp, err := d.readSensorData(BMP388_RegTemp)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// pulled from C driver: https://github.com/BoschSensortec/BMP3-Sensor-API/blob/master/bmp3.c
// 	partialData1 := rawTemp - (256 * int64(d.cali.t1))
// 	partialData2 := int64(d.cali.t2) * partialData1
// 	partialData3 := (partialData1 * partialData1)
// 	partialData4 := partialData3 * int64(d.cali.t3)
// 	partialData5 := (partialData2 * 262144) + partialData4
// 	return partialData5 / 4294967296, nil

// }

// func (d Bmp388) ReadTemperature() (int32, error) {

// 	tlin, err := d.tlinCompensate()
// 	if err != nil {
// 		return 0, err
// 	}

// 	temp := (tlin * 25) / 16384
// 	return int32(temp), nil
// }

// func (d Bmp388) ReadPressure() (int32, error) {

// 	tlin, err := d.tlinCompensate()
// 	if err != nil {
// 		return 0, err
// 	}
// 	rawPress, err := d.readSensorData(BMP388_RegPress)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// code pulled from bmp388 C driver: https://github.com/BoschSensortec/BMP3-Sensor-API/blob/master/bmp3.c
// 	partialData1 := tlin * tlin
// 	partialData2 := partialData1 / 64
// 	partialData3 := (partialData2 * tlin) / 256
// 	partialData4 := (int64(d.cali.p8) * partialData3) / 32
// 	partialData5 := (int64(d.cali.p7) * partialData1) * 16
// 	partialData6 := (int64(d.cali.p6) * tlin) * 4194304
// 	offset := (int64(d.cali.p5) * 140737488355328) + partialData4 + partialData5 + partialData6
// 	partialData2 = (int64(d.cali.p4) * partialData3) / 32
// 	partialData4 = (int64(d.cali.p3) * partialData1) * 4
// 	partialData5 = (int64(d.cali.p2) - 16384) * tlin * 2097152
// 	sensitivity := ((int64(d.cali.p1) - 16384) * 70368744177664) + partialData2 + partialData4 + partialData5
// 	partialData1 = (sensitivity / 16777216) * rawPress
// 	partialData2 = int64(d.cali.p10) * tlin
// 	partialData3 = partialData2 + (65536 * int64(d.cali.p9))
// 	partialData4 = (partialData3 * rawPress) / 8192

// 	// dividing by 10 followed by multiplying by 10
// 	// To avoid overflow caused by (pressure * partial_data4)
// 	partialData5 = (rawPress * (partialData4 / 10)) / 512
// 	partialData5 = partialData5 * 10
// 	partialData6 = (int64)(uint64(rawPress) * uint64(rawPress))
// 	partialData2 = (int64(d.cali.p11) * partialData6) / 65536
// 	partialData3 = (partialData2 * rawPress) / 128
// 	partialData4 = (offset / 4) + partialData1 + partialData5 + partialData3
// 	compPress := ((uint64(partialData4) * 25) / uint64(1099511627776))
// 	return int32(compPress), nil
// }

// func (d Bmp388) SoftReset() error {
// 	err := d.writeRegister(BMP388_RegCmd, BMP388_SoftReset)
// 	if err != nil {
// 		return errSoftReset
// 	}
// 	return nil
// }

// func (d Bmp388) SetMode(mode bmp388Mode) error {
// 	d.Config.Mode = mode
// 	return d.writeRegister(BMP388_RegPwrCtrl, BMP388_PwrPress|BMP388_PwrTemp|byte(d.Config.Mode))
// }

// func (d Bmp388) readSensorData(register byte) (data int64, err error) {

// 	if !d.Connected() {
// 		return 0, errNotConnected
// 	}

// 	// put the sensor back into forced mode to get a reading, the sensor goes back to sleep after taking one read in
// 	// forced mode
// 	if d.Config.Mode != Normal {
// 		err = d.SetMode(Forced)
// 		if err != nil {
// 			return
// 		}
// 	}

// 	bytes, err := d.readRegister(register, 3)
// 	if err != nil {
// 		return
// 	}
// 	data = int64(bytes[2])<<16 | int64(bytes[1])<<8 | int64(bytes[0])
// 	return
// }

// func (d Bmp388) readRegister(register byte, len int) (data []byte, err error) {
// 	data = make([]byte, len)
// 	err = d.bus.ReadRegister(d.Address, register, data)
// 	return
// }

// func (d Bmp388) writeRegister(register byte, data byte) error {
// 	return d.bus.WriteRegister(d.Address, register, []byte{data})
// }

// func (d Bmp388) configurationError() bool {
// 	data, err := d.readRegister(BMP388_RegErr, 1)
// 	return err == nil && (data[0]&0x04) != 0
// }
