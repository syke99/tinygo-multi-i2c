package multi

import (
	"errors"
	"math"
	"time"
)

// These variables belong to BMP388
var (
	errConfigWrite  = errors.New("bmp388: failed to configure sensor, check connection")
	errConfig       = errors.New("bmp388: there is a problem with the configuration, try reducing ODR")
	errCaliRead     = errors.New("bmp388: failed to read calibration coefficient register")
	errSoftReset    = errors.New("bmp388: failed to perform a soft reset")
	errNotConnected = errors.New("bmp388: not connected")
)

//-------------------------------------------------------------------------------------
// ADXL345
func (p adxl345PowerCtl) toByte() (bits uint8) {
	bits = 0x00
	bits = bits | (p.link << 5)
	bits = bits | (p.autoSleep << 4)
	bits = bits | (p.measure << 3)
	bits = bits | (p.sleep << 2)
	bits = bits | p.wakeUp

	return bits
}

// toByte returns a byte from the dataFormat configuration
func (d adxl345DataFormat) toByte() (bits uint8) {
	bits = 0x00
	bits = bits | (d.selfTest << 7)
	bits = bits | (d.spi << 6)
	bits = bits | (d.intInvert << 5)
	bits = bits | (d.fullRes << 3)
	bits = bits | (d.justify << 2)
	bits = bits | uint8(d.sensorRange)

	return bits
}

// toByte returns a byte from the bwRate configuration
func (b adxl345BwRate) toByte() (bits uint8) {
	bits = 0x00
	bits = bits | (b.lowPower << 4)
	bits = bits | uint8(b.rate)

	return bits
}

// Halt stops the sensor, values will not updated
func (d Adxl345) Halt() {
	d.powerCtl.measure = 0
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_POWER_CTL, []byte{d.powerCtl.toByte()})
}

// Restart makes reading the sensor working again after a halt
func (d Adxl345) Restart() {
	d.powerCtl.measure = 1
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_POWER_CTL, []byte{d.powerCtl.toByte()})
}

// ReadAcceleration reads the current acceleration from the device and returns
// it in µg (micro-gravity). When one of the axes is pointing straight to Earth
// and the sensor is not moving the returned value will be around 1000000 or
// -1000000.
func (d Adxl345) ReadAcceleration() (x int32, y int32, z int32, err error) {
	rx, ry, rz := d.ReadRawAcceleration()

	x = d.dataFormat.convertToIS(rx)
	y = d.dataFormat.convertToIS(ry)
	z = d.dataFormat.convertToIS(rz)

	return
}

// ReadRawAcceleration reads the sensor values and returns the raw x, y and z axis
// from the adxl345.
func (d Adxl345) ReadRawAcceleration() (x int32, y int32, z int32) {
	data := []byte{0, 0, 0, 0, 0, 0}
	d.bus.ReadRegister(uint8(d.Address), ADX1345_REG_DATAX0, data)

	x = readAdxl345IntLE(data[0], data[1])
	y = readAdxl345IntLE(data[2], data[3])
	z = readAdxl345IntLE(data[4], data[5])

	return
}

// UseLowPower sets the ADXL345 to use the low power mode.
func (d Adxl345) UseLowPower(power bool) {
	if power {
		d.bwRate.lowPower = 1
	} else {
		d.bwRate.lowPower = 0
	}
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_BW_RATE, []byte{d.bwRate.toByte()})
}

// SetRate change the current rate of the sensor
func (d Adxl345) SetRate(rate Adxl345Rate) bool {
	d.bwRate.rate = rate & 0x0F
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_BW_RATE, []byte{d.bwRate.toByte()})
	return true
}

// SetRange change the current range of the sensor
func (d Adxl345) SetRange(sensorRange Adxl345Range) bool {
	d.dataFormat.sensorRange = sensorRange & 0x03
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_DATA_FORMAT, []byte{d.dataFormat.toByte()})
	return true
}

// convertToIS adjusts the raw values from the adxl345 with the range configuration
func (d adxl345DataFormat) convertToIS(rawValue int32) int32 {
	switch d.sensorRange {
	case ADX1345_RANGE_2G:
		return rawValue * 4 // rawValue * 2 * 1000 / 512
	case ADX1345_RANGE_4G:
		return rawValue * 8 // rawValue * 4 * 1000 / 512
	case ADX1345_RANGE_8G:
		return rawValue * 16 // rawValue * 8 * 1000 / 512
	case ADX1345_RANGE_16G:
		return rawValue * 32 // rawValue * 16 * 1000 / 512
	default:
		return 0
	}
}

// readInt converts two bytes to int16
func readAdxl345IntLE(msb byte, lsb byte) int32 {
	return int32(uint16(msb) | uint16(lsb)<<8)
}

//-------------------------------------------------------------------------------------
// AMG88XX

// ReadPixels returns the 64 values (8x8 grid) of the sensor converted to  millicelsius
func (d Amg88xx) ReadPixels(buffer *[64]int16) {
	d.bus.ReadRegister(uint8(d.Address), AMG88XX_PIXEL_OFFSET, d.data)
	for i := 0; i < 64; i++ {
		buffer[i] = int16((uint16(d.data[2*i+1]) << 8) | uint16(d.data[2*i]))
		if (buffer[i] & (1 << 11)) > 0 { // temperature negative
			buffer[i] &= ^(1 << 11)
			buffer[i] = -buffer[i]
		}
		buffer[i] *= AMG88XX_PIXEL_TEMP_CONVERSION
	}
}

func (d Amg88xx) SetPCTL(pctl uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_PCTL, []byte{pctl})
}

func (d Amg88xx) SetReset(rst uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_RST, []byte{rst})
}

func (d Amg88xx) SetFrameRate(framerate uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_FPSC, []byte{framerate & 0x01})
}

func (d Amg88xx) SetMovingAverageMode(mode bool) {
	var value uint8
	if mode {
		value = 1
	}
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_AVE, []byte{value << 5})
}

func (d Amg88xx) SetInterruptLevels(high int16, low int16) {
	d.SetInterruptLevelsHysteresis(high, low, (high*95)/100)
}

func (d Amg88xx) SetInterruptLevelsHysteresis(high int16, low int16, hysteresis int16) {
	high = high / AMG88XX_PIXEL_TEMP_CONVERSION
	if high < -4095 {
		high = -4095
	}
	if high > 4095 {
		high = 4095
	}
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8(high & 0xFF)})
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8((high & 0xFF) >> 4)})

	low = low / AMG88XX_PIXEL_TEMP_CONVERSION
	if low < -4095 {
		low = -4095
	}
	if low > 4095 {
		low = 4095
	}
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8(low & 0xFF)})
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8((low & 0xFF) >> 4)})

	hysteresis = hysteresis / AMG88XX_PIXEL_TEMP_CONVERSION
	if hysteresis < -4095 {
		hysteresis = -4095
	}
	if hysteresis > 4095 {
		hysteresis = 4095
	}
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8(hysteresis & 0xFF)})
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8((hysteresis & 0xFF) >> 4)})
}

func (d Amg88xx) EnableInterrupt() {
	d.interruptEnable = 1
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTC, []byte{((uint8(d.interruptMode) << 1) | d.interruptEnable) & 0x03})
}

func (d Amg88xx) DisableInterrupt() {
	d.interruptEnable = 0
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTC, []byte{((uint8(d.interruptMode) << 1) | d.interruptEnable) & 0x03})
}

func (d Amg88xx) SetInterruptMode(mode Amg88xxInterruptMode) {
	d.interruptMode = mode
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTC, []byte{((uint8(d.interruptMode) << 1) | d.interruptEnable) & 0x03})
}

func (d Amg88xx) GetInterrupt() []uint8 {
	data := make([]uint8, 8)
	d.bus.ReadRegister(uint8(d.Address), AMG88XX_INT_OFFSET, data)
	return data
}

func (d Amg88xx) ClearInterrupt() {
	d.SetReset(AMG88XX_FLAG_RESET)
}

func (d Amg88xx) ReadThermistor() int16 {
	data := make([]uint8, 2)
	d.bus.ReadRegister(uint8(d.Address), AMG88XX_TTHL, data)
	return (int16((uint16(data[1])<<8)|uint16(data[0])) * AMG88XX_THERMISTOR_CONVERSION) / 10
}

//-------------------------------------------------------------------------------------
// AT24CX
// skip for now

//-------------------------------------------------------------------------------------
// BH1750
func (d Bh1750) RawSensorData() uint16 {

	buf := []byte{1, 0}
	d.bus.Tx(d.Address, nil, buf)
	return (uint16(buf[0]) << 8) | uint16(buf[1])
}

func (d Bh1750) Illuminance() int32 {

	lux := uint32(d.RawSensorData())
	var coef uint32
	if d.mode == BH1750_CONTINUOUS_HIGH_RES_MODE || d.mode == BH1750_ONE_TIME_HIGH_RES_MODE {
		coef = BH1750_HIGH_RES
	} else if d.mode == BH1750_CONTINUOUS_HIGH_RES_MODE_2 || d.mode == BH1750_ONE_TIME_HIGH_RES_MODE_2 {
		coef = BH1750_HIGH_RES2
	} else {
		coef = BH1750_LOW_RES
	}
	// 100 * coef * lux * (5/6)
	// 5/6 = measurement accuracy as per the datasheet
	return int32(250 * coef * lux / 3)
}

func (d Bh1750) SetMode(mode byte) {
	d.mode = mode
	d.bus.Tx(d.Address, []byte{byte(d.mode)}, nil)
	time.Sleep(10 * time.Millisecond)
}

//-------------------------------------------------------------------------------------
// BLINKM
func (d Blinkm) Version() (major, minor byte, err error) {
	version := []byte{0, 0}
	d.bus.Tx(d.Address, []byte{BLINKM_GET_FIRMWARE}, version)
	return version[0], version[1], nil
}

func (d Blinkm) SetRGB(r, g, b byte) error {
	d.bus.Tx(d.Address, []byte{BLINKM_TO_RGB, r, g, b}, nil)
	return nil
}

func (d Blinkm) GetRGB() (r, g, b byte, err error) {
	color := []byte{0, 0, 0}
	d.bus.Tx(d.Address, []byte{BLINKM_GET_RGB}, color)
	return color[0], color[1], color[2], nil
}

func (d Blinkm) FadeToRGB(r, g, b byte) error {
	d.bus.Tx(d.Address, []byte{BLINKM_FADE_TO_RGB, r, g, b}, nil)
	return nil
}

func (d Blinkm) StopScript() error {
	d.bus.Tx(d.Address, []byte{BLINKM_STOP_SCRIPT}, nil)
	return nil
}

//-------------------------------------------------------------------------------------
// BME280
func (d Bme280) Reset() {
	d.bus.WriteRegister(uint8(d.Address), BME280_CMD_RESET, []byte{0xB6})
}

func (d Bme280) ReadTemperature() (int32, error) {
	data, err := d.readData()
	if err != nil {
		return 0, err
	}

	temp, _ := d.calculateTemp(data)
	return temp, nil
}

func (d Bme280) ReadPressure() (int32, error) {
	data, err := d.readData()
	if err != nil {
		return 0, err
	}
	_, tFine := d.calculateTemp(data)
	pressure := d.calculatePressure(data, tFine)
	return pressure, nil
}

func (d Bme280) ReadHumidity() (int32, error) {
	data, err := d.readData()
	if err != nil {
		return 0, err
	}
	_, tFine := d.calculateTemp(data)
	humidity := d.calculateHumidity(data, tFine)
	return humidity, nil
}

func (d Bme280) ReadAltitude() (alt int32, err error) {
	mPa, _ := d.ReadPressure()
	atmP := float32(mPa) / 100000
	alt = int32(44330.0 * (1.0 - math.Pow(float64(atmP/BME280_SEALEVEL_PRESSURE), 0.1903)))
	return
}

func (d Bme280) calculateTemp(data [8]byte) (int32, int32) {

	rawTemp := convert3Bytes(data[3], data[4], data[5])

	var1 := (((rawTemp >> 3) - (int32(d.calibrationCoefficients.t1) << 1)) * int32(d.calibrationCoefficients.t2)) >> 11
	var2 := (((((rawTemp >> 4) - int32(d.calibrationCoefficients.t1)) * ((rawTemp >> 4) - int32(d.calibrationCoefficients.t1))) >> 12) * int32(d.calibrationCoefficients.t3)) >> 14

	tFine := var1 + var2
	T := (tFine*5 + 128) >> 8
	return (10 * T), tFine
}

func (d Bme280) readData() (data [8]byte, err error) {
	err = d.bus.ReadRegister(uint8(d.Address), BME280_REG_PRESSURE, data[:])
	if err != nil {
		println(err)
		return
	}
	return
}

func (d Bme280) calculatePressure(data [8]byte, tFine int32) int32 {

	rawPressure := convert3Bytes(data[0], data[1], data[2])

	var1 := int64(tFine) - 128000
	var2 := var1 * var1 * int64(d.calibrationCoefficients.p6)
	var2 = var2 + ((var1 * int64(d.calibrationCoefficients.p5)) << 17)
	var2 = var2 + (int64(d.calibrationCoefficients.p4) << 35)
	var1 = ((var1 * var1 * int64(d.calibrationCoefficients.p3)) >> 8) + ((var1 * int64(d.calibrationCoefficients.p2)) << 12)
	var1 = ((int64(1) << 47) + var1) * int64(d.calibrationCoefficients.p1) >> 33

	if var1 == 0 {
		return 0 // avoid exception caused by division by zero
	}
	p := int64(1048576 - rawPressure)
	p = (((p << 31) - var2) * 3125) / var1
	var1 = (int64(d.calibrationCoefficients.p9) * (p >> 13) * (p >> 13)) >> 25
	var2 = (int64(d.calibrationCoefficients.p8) * p) >> 19

	p = ((p + var1 + var2) >> 8) + (int64(d.calibrationCoefficients.p7) << 4)
	p = (p / 256)
	return int32(1000 * p)
}

func (d Bme280) calculateHumidity(data [8]byte, tFine int32) int32 {

	rawHumidity := convert2Bytes(data[6], data[7])

	h := float32(tFine) - 76800

	if h == 0 {
		println("invalid value")
	}

	var1 := float32(rawHumidity) - (float32(d.calibrationCoefficients.h4)*64.0 +
		(float32(d.calibrationCoefficients.h5) / 16384.0 * h))

	var2 := float32(d.calibrationCoefficients.h2) / 65536.0 *
		(1.0 + float32(d.calibrationCoefficients.h6)/67108864.0*h*
			(1.0+float32(d.calibrationCoefficients.h3)/67108864.0*h))

	h = var1 * var2
	h = h * (1 - float32(d.calibrationCoefficients.h1)*h/524288)
	return int32(100 * h)

}

func convert2Bytes(msb byte, lsb byte) int32 {
	return int32(readUint(msb, lsb))
}

func convert3Bytes(msb byte, b1 byte, lsb byte) int32 {
	return int32(((((uint32(msb) << 8) | uint32(b1)) << 8) | uint32(lsb)) >> 4)
}

//-------------------------------------------------------------------------------------
// BMP280
func (d Bmp280) Reset() {
	d.bus.WriteRegister(uint8(d.Address), BMP280_REG_RESET, []byte{BMP280_CMD_RESET})
}

func (d Bmp280) ReadTemperature() (temperature int32, err error) {
	data, err := d.readData(BMP280_REG_TEMP, 3)
	if err != nil {
		return
	}

	rawTemp := convert3Bytes(data[0], data[1], data[2])

	// Datasheet: 8.2 Compensation formula in 32 bit fixed point
	// Temperature compensation
	var1 := ((rawTemp >> 3) - int32(d.cali.t1<<1)) * int32(d.cali.t2) >> 11
	var2 := (((rawTemp >> 4) - int32(d.cali.t1)) * ((rawTemp >> 4) - int32(d.cali.t1)) >> 12) *
		int32(d.cali.t3) >> 14

	tFine := var1 + var2

	// Convert from degrees to milli degrees by multiplying by 10.
	// Will output 30250 milli degrees celsius for 30.25 degrees celsius
	temperature = 10 * ((tFine*5 + 128) >> 8)
	return
}

func (d Bmp280) ReadPressure() (pressure int32, err error) {
	// First 3 bytes are Pressure, last 3 bytes are Temperature
	data, err := d.readData(BMP280_REG_PRES, 6)
	if err != nil {
		return
	}

	rawTemp := convert3Bytes(data[3], data[4], data[5])

	// Datasheet: 8.2 Compensation formula in 32 bit fixed point
	// Calculate tFine (temperature), used for the Pressure compensation
	var1 := ((rawTemp >> 3) - int32(d.cali.t1<<1)) * int32(d.cali.t2) >> 11
	var2 := (((rawTemp >> 4) - int32(d.cali.t1)) * ((rawTemp >> 4) - int32(d.cali.t1)) >> 12) *
		int32(d.cali.t3) >> 14

	tFine := var1 + var2

	rawPres := convert3Bytes(data[0], data[1], data[2])

	// Datasheet: 8.2 Compensation formula in 32 bit fixed point
	// Pressure compensation
	var1 = (tFine >> 1) - 64000
	var2 = (((var1 >> 2) * (var1 >> 2)) >> 11) * int32(d.cali.p6)
	var2 = var2 + ((var1 * int32(d.cali.p5)) << 1)
	var2 = (var2 >> 2) + (int32(d.cali.p4) << 16)
	var1 = (((int32(d.cali.p3) * (((var1 >> 2) * (var1 >> 2)) >> 13)) >> 3) +
		((int32(d.cali.p2) * var1) >> 1)) >> 18
	var1 = ((32768 + var1) * int32(d.cali.p1)) >> 15

	if var1 == 0 {
		return 0, nil
	}

	p := uint32(((1048576 - rawPres) - (var2 >> 12)) * 3125)
	if p < 0x80000000 {
		p = (p << 1) / uint32(var1)
	} else {
		p = (p / uint32(var1)) * 2
	}

	var1 = (int32(d.cali.p9) * int32(((p>>3)*(p>>3))>>13)) >> 12
	var2 = (int32(p>>2) * int32(d.cali.p8)) >> 13

	return 1000 * (int32(p) + ((var1 + var2 + int32(d.cali.p7)) >> 4)), nil
}

func (d Bmp280) readData(register int, n int) ([]byte, error) {
	// If not in normal mode, set the mode to FORCED mode, to prevent incorrect measurements
	// After the measurement in FORCED mode, the sensor will return to SLEEP mode
	if d.Mode != BMP280_MODE_NORMAL {
		config := uint(d.Temperature<<5) | uint(d.Pressure<<2) | uint(BMP280_MODE_FORCED)
		d.bus.WriteRegister(uint8(d.Address), BMP280_REG_CTRL_MEAS, []byte{byte(config)})
	}

	// Check STATUS register, wait if data is not available yet
	status := make([]byte, 1)
	for d.bus.ReadRegister(uint8(d.Address), uint8(BMP280_REG_STATUS), status[0:]); status[0] != 4 && status[0] != 0; d.bus.ReadRegister(uint8(d.Address), uint8(BMP280_REG_STATUS), status[0:]) {
		time.Sleep(time.Millisecond)
	}

	// Read the requested register
	data := make([]byte, n)
	err := d.bus.ReadRegister(uint8(d.Address), uint8(register), data[:])
	return data, err
}

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

//-------------------------------------------------------------------------------------
// VL53L1x
func (d Vl53l1x) Connected() bool {
	return d.readReg16Bit(VL53L1x_WHO_AM_I) == VL53L1x_CHIP_ID
}

// writeReg sends a single byte to the specified register address
func (d Vl53l1x) writeReg(reg uint16, value uint8) {
	msb := byte((reg >> 8) & 0xFF)
	lsb := byte(reg & 0xFF)
	d.bus.Tx(d.Address, []byte{msb, lsb, value}, nil)
}

// writeReg16Bit sends two bytes to the specified register address
func (d Vl53l1x) writeReg16Bit(reg uint16, value uint16) {
	data := make([]byte, 4)
	data[0] = byte((reg >> 8) & 0xFF)
	data[1] = byte(reg & 0xFF)
	data[2] = byte((value >> 8) & 0xFF)
	data[3] = byte(value & 0xFF)
	d.bus.Tx(d.Address, data, nil)
}

// writeReg32Bit sends four bytes to the specified register address
func (d Vl53l1x) writeReg32Bit(reg uint16, value uint32) {
	data := make([]byte, 6)
	data[0] = byte((reg >> 8) & 0xFF)
	data[1] = byte(reg & 0xFF)
	data[2] = byte((value >> 24) & 0xFF)
	data[3] = byte((value >> 16) & 0xFF)
	data[4] = byte((value >> 8) & 0xFF)
	data[5] = byte(value & 0xFF)
	d.bus.Tx(d.Address, data, nil)
}

// readReg reads a single byte from the specified address
func (d Vl53l1x) readReg(reg uint16) uint8 {
	data := []byte{0}
	msb := byte((reg >> 8) & 0xFF)
	lsb := byte(reg & 0xFF)
	d.bus.Tx(d.Address, []byte{msb, lsb}, data)
	return data[0]
}

// readReg16Bit reads two bytes from the specified address
// and returns it as a uint16
func (d Vl53l1x) readReg16Bit(reg uint16) uint16 {
	data := []byte{0, 0}
	msb := byte((reg >> 8) & 0xFF)
	lsb := byte(reg & 0xFF)
	d.bus.Tx(d.Address, []byte{msb, lsb}, data)
	return readUint(data[0], data[1])
}

// readReg32Bit reads four bytes from the specified address
// and returns it as a uint32
func (d Vl53l1x) readReg32Bit(reg uint16) uint32 {
	data := make([]byte, 4)
	msb := byte((reg >> 8) & 0xFF)
	lsb := byte(reg & 0xFF)
	d.bus.Tx(d.Address, []byte{msb, lsb}, data)
	return readUint32(data)
}

func (d Vl53l1x) SetMeasurementTimingBudget(budgetMicroseconds uint32) bool {
	if budgetMicroseconds <= VL53L1x_TIMING_GUARD {
		return false
	}
	budgetMicroseconds -= VL53L1x_TIMING_GUARD
	if budgetMicroseconds > 1100000 {
		return false
	}
	rangeConfigTimeout := budgetMicroseconds / 2
	// Update Macro Period for Range A VCSEL Period
	macroPeriod := d.calculateMacroPeriod(uint32(d.readReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A)))

	// Update Phase timeout - uses Timing A
	phasecalTimeoutMclks := timeoutMicrosecondsToMclks(1000, macroPeriod)
	if phasecalTimeoutMclks > 0xFF {
		phasecalTimeoutMclks = 0xFF
	}
	d.writeReg(VL53L1x_PHASECAL_CONFIG_TIMEOUT_MACROP, uint8(phasecalTimeoutMclks))

	// Update MM Timing A timeout
	d.writeReg16Bit(VL53L1x_MM_CONFIG_TIMEOUT_MACROP_A, encodeTimeout(timeoutMicrosecondsToMclks(1, macroPeriod)))
	// Update Range Timing A timeout
	d.writeReg16Bit(VL53L1x_RANGE_CONFIG_TIMEOUT_MACROP_A, encodeTimeout(timeoutMicrosecondsToMclks(rangeConfigTimeout, macroPeriod)))

	macroPeriod = d.calculateMacroPeriod(uint32(d.readReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A)))
	// Update MM Timing B timeout
	d.writeReg16Bit(VL53L1x_MM_CONFIG_TIMEOUT_MACROP_B, encodeTimeout(timeoutMicrosecondsToMclks(1, macroPeriod)))
	// Update Range Timing B timeout
	d.writeReg16Bit(VL53L1x_RANGE_CONFIG_TIMEOUT_MACROP_B, encodeTimeout(timeoutMicrosecondsToMclks(rangeConfigTimeout, macroPeriod)))

	return true
}

func (d Vl53l1x) calculateMacroPeriod(vcselPeriod uint32) uint32 {
	pplPeriodMicroseconds := (uint32(1) << 30) / uint32(d.fastOscillatorFreq)
	vcselPeriodPclks := (vcselPeriod + 1) << 1
	macroPeriodMicroseconds := 2304 * pplPeriodMicroseconds
	macroPeriodMicroseconds >>= 6
	macroPeriodMicroseconds *= vcselPeriodPclks
	macroPeriodMicroseconds >>= 6
	return macroPeriodMicroseconds
}

func (d Vl53l1x) SetDistanceMode(mode vl53l1xDistanceMode) bool {
	budgetMicroseconds := d.GetMeasurementTimingBudget()
	switch mode {
	case VL53L1x_SHORT:
		// timing config
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A, 0x07)
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_B, 0x05)
		d.writeReg(VL53L1x_RANGE_CONFIG_VALID_PHASE_HIGH, 0x38)

		// dynamic config
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD0, 0x07)
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD1, 0x05)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD0, 6)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD1, 6)
		break
	case VL53L1x_MEDIUM:
		// timing config
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A, 0x0B)
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_B, 0x09)
		d.writeReg(VL53L1x_RANGE_CONFIG_VALID_PHASE_HIGH, 0x78)

		// dynamic config
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD0, 0x0B)
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD1, 0x09)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD0, 10)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD1, 10)
		break
	case VL53L1x_LONG:
		// timing config
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A, 0x0F)
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_B, 0x0D)
		d.writeReg(VL53L1x_RANGE_CONFIG_VALID_PHASE_HIGH, 0xB8)

		// dynamic config
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD0, 0x0F)
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD1, 0x0D)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD0, 14)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD1, 14)
		break
	default:
		return false
	}

	d.SetMeasurementTimingBudget(budgetMicroseconds)
	d.mode = mode
	return true
}

func (d Vl53l1x) GetMeasurementTimingBudget() uint32 {
	macroPeriod := d.calculateMacroPeriod(uint32(d.readReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A)))
	rangeConfigTimeout := timeoutMclksToMicroseconds(decodeTimeout(d.readReg16Bit(VL53L1x_RANGE_CONFIG_TIMEOUT_MACROP_A)), macroPeriod)
	return 2 * uint32(rangeConfigTimeout) * VL53L1x_TIMING_GUARD
}
