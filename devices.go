package multi

import (
	"tinygo.org/x/drivers"
)

// Device structs for creating new devices
// Capitalized structs are devices, lower-case
// structs and non-stuct types are requried by
// the following capitalized struct

type Range uint8
type Rate uint8

// Internal structure for the power configuration
type powerCtl struct {
	link      uint8
	autoSleep uint8
	measure   uint8
	sleep     uint8
	wakeUp    uint8
}

// Internal structure for the sensor's data format configuration
type dataFormat struct {
	selfTest    uint8
	spi         uint8
	intInvert   uint8
	fullRes     uint8
	justify     uint8
	sensorRange Range
}

// Internal structure for the sampling rate configuration
type bwRate struct {
	lowPower uint8
	rate     Rate
}

type Adxl345 struct {
	bus        drivers.I2C
	Address    uint16
	powerCtl   powerCtl
	dataFormat dataFormat
	bwRate     bwRate
}

type Aht20 struct {
	bus      drivers.I2C
	Address  uint16
	humidity uint32
	temp     uint32
}

type InterruptMode uint8

type Amg88xx struct {
	bus             drivers.I2C
	Address         uint16
	data            []uint8
	interruptMode   InterruptMode
	interruptEnable uint8
}

type gestureData struct {
	detected    uint8
	threshold   uint8
	sensitivity uint8
	gXDelta     int16
	gYDelta     int16
	gXPrevDelta int16
	gYPrevDelta int16
	received    bool
}

type Apds9960 struct {
	bus     drivers.I2C
	Address uint8
	mode    uint8
	gesture gestureData
}

type At24cx struct {
	bus               drivers.I2C
	Address           uint16
	pageSize          uint16
	currentRAMAddress uint16
	startRAMAddress   uint16
	endRAMAddress     uint16
}

type Bh1750 struct {
	bus     drivers.I2C
	Address uint16
	mode    byte
}

type Blinkm struct {
	bus     drivers.I2C
	Address uint16
}

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

type bmp180CalibrationCoefficients struct {
	ac1 int16
	ac2 int16
	ac3 int16
	ac4 uint16
	ac5 uint16
	ac6 uint16
	b1  int16
	b2  int16
	mb  int16
	mc  int16
	md  int16
}

type bmp180OversamplingMode uint

type Bmp180 struct {
	bus                     drivers.I2C
	Address                 uint16
	mode                    bmp180OversamplingMode
	calibrationCoefficients bmp180CalibrationCoefficients
}

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

type bmpOversampling byte
type bmpMode byte
type bmpOutputDataRate byte
type bmpFilterCoefficient byte

type bmp388Config struct {
	Pressure    byte
	Temperature bmpOversampling
	Mode        bmpMode
	ODR         bmpOutputDataRate
	IIR         bmpFilterCoefficient
}

type Bmp388 struct {
	bus     drivers.I2C
	Address uint8
	cali    bmp388CalibrationCoefficients
	Config  bmp388Config
}

type Ds3231 struct {
	bus     drivers.I2C
	Address uint16
}

type Hts221 struct {
	bus              drivers.I2C
	Address          uint8
	humiditySlope    float32
	humidityZero     float32
	temperatureSlope float32
	temperatureZero  float32
}

type Ina260 struct {
	bus     drivers.I2C
	Address uint16
}

type Lis2mdl struct {
	bus        drivers.I2C
	Address    uint8
	PowerMode  uint8
	SystemMode uint8
	DataRate   uint8
}

type Lis3dh struct {
	bus     drivers.I2C
	Address uint16
	r       Range
}

type Lps22hb struct {
	bus     drivers.I2C
	Address uint8
}

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

type Mag3110 struct {
	bus     drivers.I2C
	Address uint16
}

type mma8653Sensitivity uint8

type Mma8653 struct {
	bus         drivers.I2C
	Address     uint16
	sensitivity mma8653Sensitivity
}

type Mpu6050 struct {
	bus     drivers.I2C
	Address uint16
}

type Pcf8563 struct {
	bus     drivers.I2C
	Address uint16
}

type Sht3x struct {
	bus     drivers.I2C
	Address uint16
}

type Tmp102 struct {
	bus     drivers.I2C
	address uint8
}

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
