package multi

import (
	"tinygo.org/x/drivers"
)

// Device structs for creating new devices
// Capitalized structs are devices, lower-case
// structs and non-stuct types are requried by
// the following capitalized struct

type Adxl345Range uint8
type Adxl345Rate uint8

// Internal structure for the power configuration
type powerCtl struct {
	link      uint8
	autoSleep uint8
	measure   uint8
	sleep     uint8
	wakeUp    uint8
}

// Internal structure for the sensor's data format configuration
type adxl345DataFormat struct {
	selfTest    uint8
	spi         uint8
	intInvert   uint8
	fullRes     uint8
	justify     uint8
	sensorRange Adxl345Range
}

// Internal structure for the sampling rate configuration
type adxl345BwRate struct {
	lowPower uint8
	rate     Adxl345Rate
}

type Adxl345 struct {
	bus        drivers.I2C
	Address    uint16
	powerCtl   powerCtl
	dataFormat adxl345DataFormat
	bwRate     adxl345BwRate
}

// amg88xx
type Amg88xxInterruptMode uint8

type Amg88xx struct {
	bus             drivers.I2C
	Address         uint16
	data            []uint8
	interruptMode   Amg88xxInterruptMode
	interruptEnable uint8
}

// at24cx
type At24cx struct {
	bus               drivers.I2C
	Address           uint16
	pageSize          uint16
	currentRAMAddress uint16
	startRAMAddress   uint16
	endRAMAddress     uint16
}

// bh1740
type Bh1750 struct {
	bus     drivers.I2C
	Address uint16
	mode    byte
}

// blinkm
type Blinkm struct {
	bus     drivers.I2C
	Address uint16
}

// bme280
type bme280CalibrationCoefficients struct {
	t1 uint16
	t2 int16
	t3 int16
	p1 uint16
	p2 int16
	p3 int16
	p4 int16
	p5 int16
	p6 int16
	p7 int16
	p8 int16
	p9 int16
	h1 uint8
	h2 int16
	h3 uint8
	h4 int16
	h5 int16
	h6 int8
}

type Bme280 struct {
	bus                     drivers.I2C
	Address                 uint16
	calibrationCoefficients bme280CalibrationCoefficients
}

// bpm280
type bmp280CalibrationCoefficients struct {
	// Temperature compensation
	t1 uint16
	t2 int16
	t3 int16

	// Pressure compensation
	p1 uint16
	p2 int16
	p3 int16
	p4 int16
	p5 int16
	p6 int16
	p7 int16
	p8 int16
	p9 int16
}

type bmp280Oversampling uint
type bmp280Mode uint
type bmp290Standby uint
type bmp280Filter uint

type Bmp280 struct {
	bus         drivers.I2C
	Address     uint16
	cali        bmp280CalibrationCoefficients
	Temperature bmp280Oversampling
	Pressure    bmp280Oversampling
	Mode        bmp280Mode
	Standby     bmp290Standby
	Filter      bmp280Filter
}

// bmp388
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

// ds3231
type Ds3231Mode uint8

type Ds3231 struct {
	bus     drivers.I2C
	Address uint16
}

// ina260
type Ina260 struct {
	bus     drivers.I2C
	Address uint16
}

// lis3dh
type Lis3dh struct {
	bus     drivers.I2C
	Address uint16
	r       Lis3dhRange
}

// lps22hb
type Lps22hb struct {
	bus     drivers.I2C
	Address uint8
}

// lsm6ds3
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

// mpu6050
type Mpu6050 struct {
	bus     drivers.I2C
	Address uint16
}

// sht3x
type Sht3x struct {
	bus     drivers.I2C
	Address uint16
}

// vl53l1x
type vl53l1xDistanceMode uint8
type vl53l1xRangeStatus uint8

type vl53l1xRangingData struct {
	mm              uint16
	status          vl53l1xRangeStatus
	signalRateMCPS  int32
	ambientRateMCPS int32
}

type vl53l1xResultBuffer struct {
	status                     uint8
	streamCount                uint8
	effectiveSPADCount         uint16
	ambientRateMCPSSD0         uint16
	mmCrosstalkSD0             uint16
	signalRateCrosstalkMCPSSD0 uint16
}

type Vl53l1x struct {
	bus                drivers.I2C
	Address            uint16
	mode               vl53l1xDistanceMode
	timeout            uint32
	fastOscillatorFreq uint16
	oscillatorOffset   uint16
	calibrated         bool
	VHVInit            uint8
	VHVTimeout         uint8
	rangingData        vl53l1xRangingData
	results            vl53l1xResultBuffer
}
