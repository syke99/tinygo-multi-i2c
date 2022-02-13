package multi

// ADX1345 registers
const Adx1345AddressLow = 0x53
const Adx1345AddressHigh = 0x1D

const (
	// Data rate
	ADX1345_RATE_3200HZ Adxl345Rate = 0x0F // 3200 Hz
	ADX1345_RATE_1600HZ Adxl345Rate = 0x0E // 1600 Hz
	ADX1345_RATE_800HZ  Adxl345Rate = 0x0D // 800 Hz
	ADX1345_RATE_400HZ  Adxl345Rate = 0x0C // 400 Hz
	ADX1345_RATE_200HZ  Adxl345Rate = 0x0B // 200 Hz
	ADX1345_RATE_100HZ  Adxl345Rate = 0x0A // 100 Hz
	ADX1345_RATE_50HZ   Adxl345Rate = 0x09 // 50 Hz
	ADX1345_RATE_25HZ   Adxl345Rate = 0x08 // 25 Hz
	ADX1345_RATE_12_5HZ Adxl345Rate = 0x07 // 12.5 Hz
	ADX1345_RATE_6_25HZ Adxl345Rate = 0x06 // 6.25 Hz
	ADX1345_RATE_3_13HZ Adxl345Rate = 0x05 // 3.13 Hz
	ADX1345_RATE_1_56HZ Adxl345Rate = 0x04 // 1.56 Hz
	ADX1345_RATE_0_78HZ Adxl345Rate = 0x03 // 0.78 Hz
	ADX1345_RATE_0_39HZ Adxl345Rate = 0x02 // 0.39 Hz
	ADX1345_RATE_0_20HZ Adxl345Rate = 0x01 // 0.20 Hz
	ADX1345_RATE_0_10HZ Adxl345Rate = 0x00 // 0.10 Hz

	// Data range
	ADX1345_RANGE_2G  Adxl345Range = 0x00 // +-2 g
	ADX1345_RANGE_4G  Adxl345Range = 0x01 // +-4 g
	ADX1345_RANGE_8G  Adxl345Range = 0x02 // +-8 g
	ADX1345_RANGE_16G Adxl345Range = 0x03 // +-16 g)

	ADX1345_REG_DEVID          = 0x00 // R,     11100101,   Device ID
	ADX1345_REG_THRESH_TAP     = 0x1D // R/W,   00000000,   Tap threshold
	ADX1345_REG_OFSX           = 0x1E // R/W,   00000000,   X-axis offset
	ADX1345_REG_OFSY           = 0x1F // R/W,   00000000,   Y-axis offset
	ADX1345_REG_OFSZ           = 0x20 // R/W,   00000000,   Z-axis offset
	ADX1345_REG_DUR            = 0x21 // R/W,   00000000,   Tap duration
	ADX1345_REG_LATENT         = 0x22 // R/W,   00000000,   Tap latency
	ADX1345_REG_WINDOW         = 0x23 // R/W,   00000000,   Tap window
	ADX1345_REG_THRESH_ACT     = 0x24 // R/W,   00000000,   Activity threshold
	ADX1345_REG_THRESH_INACT   = 0x25 // R/W,   00000000,   Inactivity threshold
	ADX1345_REG_TIME_INACT     = 0x26 // R/W,   00000000,   Inactivity time
	ADX1345_REG_ACT_INACT_CTL  = 0x27 // R/W,   00000000,   Axis enable control for activity and inactiv ity detection
	ADX1345_REG_THRESH_FF      = 0x28 // R/W,   00000000,   Free-fall threshold
	ADX1345_REG_TIME_FF        = 0x29 // R/W,   00000000,   Free-fall time
	ADX1345_REG_TAP_AXES       = 0x2A // R/W,   00000000,   Axis control for single tap/double tap
	ADX1345_REG_ACT_TAP_STATUS = 0x2B // R,     00000000,   Source of single tap/double tap
	ADX1345_REG_BW_RATE        = 0x2C // R/W,   00001010,   Data rate and power mode control
	ADX1345_REG_POWER_CTL      = 0x2D // R/W,   00000000,   Power-saving features control
	ADX1345_REG_INT_ENABLE     = 0x2E // R/W,   00000000,   Interrupt enable control
	ADX1345_REG_INT_MAP        = 0x2F // R/W,   00000000,   Interrupt mapping control
	ADX1345_REG_INT_SOUCE      = 0x30 // R,     00000010,   Source of interrupts
	ADX1345_REG_DATA_FORMAT    = 0x31 // R/W,   00000000,   Data format control
	ADX1345_REG_DATAX0         = 0x32 // R,     00000000,   X-Axis Data 0
	ADX1345_REG_DATAX1         = 0x33 // R,     00000000,   X-Axis Data 1
	ADX1345_REG_DATAY0         = 0x34 // R,     00000000,   Y-Axis Data 0
	ADX1345_REG_DATAY1         = 0x35 // R,     00000000,   Y-Axis Data 1
	ADX1345_REG_DATAZ0         = 0x36 // R,     00000000,   Z-Axis Data 0
	ADX1345_REG_DATAZ1         = 0x37 // R,     00000000,   Z-Axis Data 1
	ADX1345_REG_FIFO_CTL       = 0x38 // R/W,   00000000,   FIFO control
	ADX1345_REG_FIFO_STATUS    = 0x39 // R,     00000000,   FIFO status
)

// AMG88XX registers
const Amg88xxAddressHigh = 0x69
const Amg88xxAddressLow = 0x68

const (
	AMG88XX_PCTL         = 0x00
	AMG88XX_RST          = 0x01
	AMG88XX_FPSC         = 0x02
	AMG88XX_INTC         = 0x03
	AMG88XX_STAT         = 0x04
	AMG88XX_SCLR         = 0x05
	AMG88XX_AVE          = 0x07
	AMG88XX_INTHL        = 0x08
	AMG88XX_INTHH        = 0x09
	AMG88XX_INTLL        = 0x0A
	AMG88XX_INTLH        = 0x0B
	AMG88XX_IHYSL        = 0x0C
	AMG88XX_IHYSH        = 0x0D
	AMG88XX_TTHL         = 0x0E
	AMG88XX_TTHH         = 0x0F
	AMG88XX_INT_OFFSET   = 0x010
	AMG88XX_PIXEL_OFFSET = 0x80

	// power modes
	AMG88XX_NORMAL_MODE = 0x00
	AMG88XX_SLEEP_MODE  = 0x01
	AMG88XX_STAND_BY_60 = 0x20
	AMG88XX_STAND_BY_10 = 0x21

	// resets
	AMG88XX_FLAG_RESET    = 0x30
	AMG88XX_INITIAL_RESET = 0x3F

	// frame rates
	AMG88XX_FPS_10 = 0x00
	AMG88XX_FPS_1  = 0x01

	// interrupt modes
	AMG88XX_DIFFERENCE     Amg88xxInterruptMode = 0x00
	AMG88XX_ABSOLUTE_VALUE Amg88xxInterruptMode = 0x01

	AMG88XX_PIXEL_TEMP_CONVERSION = 250
	AMG88XX_THERMISTOR_CONVERSION = 625
)

// AT24CX registers
const At24cxAddress = 0x57

// BH1750 registers
const Bh1750Address = 0x23

const (
	BH1750_POWER_DOWN                      = 0x00
	BH1750_POWER_ON                        = 0x01
	BH1750_RESET                           = 0x07
	BH1750_CONTINUOUS_HIGH_RES_MODE   byte = 0x10
	BH1750_CONTINUOUS_HIGH_RES_MODE_2 byte = 0x11
	BH1750_CONTINUOUS_LOW_RES_MODE    byte = 0x13
	BH1750_ONE_TIME_HIGH_RES_MODE     byte = 0x20
	BH1750_ONE_TIME_HIGH_RES_MODE_2   byte = 0x21
	BH1750_ONE_TIME_LOW_RES_MODE      byte = 0x23

	// resolution in 10*lx
	BH1750_HIGH_RES  = 10
	BH1750_HIGH_RES2 = 5
	BH1750_LOW_RES   = 40
)

// BLINKM registers
const BlinkmAddress = 0x09

const (
	BLINKM_TO_RGB            = 0x6e
	BLINKM_FADE_TO_RGB       = 0x63
	BLINKM_FADE_TO_HSB       = 0x68
	BLINKM_FADE_TO_RND_RGB   = 0x43
	BLINKM_FADE_TO_RND_HSB   = 0x48
	BLINKM_PLAY_LIGHT_SCRIPT = 0x70
	BLINKM_STOP_SCRIPT       = 0x6f
	BLINKM_SET_FADE          = 0x66
	BLINKM_SET_TIME          = 0x74
	BLINKM_GET_RGB           = 0x67
	BLINKM_GET_ADDRESS       = 0x61
	BLINKM_SET_ADDRESS       = 0x41
	BLINKM_GET_FIRMWARE      = 0x5a
)

// BME280 registers
const Bme280Address = 0x76

const (
	BME280_CTRL_MEAS_ADDR        = 0xF4
	BME280_CTRL_HUMIDITY_ADDR    = 0xF2
	BME280_CTRL_CONFIG           = 0xF5
	BME280_REG_PRESSURE          = 0xF7
	BME280_REG_CALIBRATION       = 0x88
	BME280_REG_CALIBRATION_H1    = 0xA1
	BME280_REG_CALIBRATION_H2LSB = 0xE1
	BME280_CMD_RESET             = 0xE0

	BME280_WHO_AM_I = 0xD0
	BME280_CHIP_ID  = 0x60
)

const (
	BME280_SEALEVEL_PRESSURE float32 = 1013.25 // in hPa
)

// BMP280 registers
const Bmp280Address = 0x77

const (
	BMP280_REG_ID        = 0xD0 // WHO_AM_I
	BMP280_REG_RESET     = 0xE0
	BMP280_REG_STATUS    = 0xF3
	BMP280_REG_CTRL_MEAS = 0xF4
	BMP280_REG_CONFIG    = 0xF5
	BMP280_REG_TEMP      = 0xFA
	BMP280_REG_PRES      = 0xF7
	BMP280_REG_CALI      = 0x88

	BMP280_CHIP_ID   = 0x58
	BMP280_CMD_RESET = 0xB6
)

const (
	BMP280_SAMPLING_SKIPPED bmp280Oversampling = iota
	BMP280_SAMPLING_1X
	BMP280_SAMPLING_2X
	BMP280_SAMPLING_4X
	BMP280_SAMPLING_8X
	BMP280_SAMPLING_16X
)

const (
	BMP280_MODE_SLEEP  bmp280Mode = 0x00
	BMP280_MODE_FORCED bmp280Mode = 0x01
	BMP280_MODE_NORMAL bmp280Mode = 0x03
)

const (
	BMP280_STANDBY_1MS bmp290Standby = iota
	BMP280_STANDBY_63MS
	BMP280_STANDBY_125MS
	BMP280_STANDBY_250MS
	BMP280_STANDBY_500MS
	BMP280_STANDBY_1000MS
	BMP280_STANDBY_2000MS
	BMP280_STANDBY_4000MS
)

const (
	BMP280_FILTER_OFF bmp280Filter = iota
	BMP280_FILTER_2X
	BMP280_FILTER_4X
	BMP280_FILTER_8X
	BMP280_FILTER_16X
)

// BMP388 registers
const Bmp388Address byte = 0x77

const (
	BMP388_RegChipId  byte = 0x00 // useful for checking the connection
	BMP388_RegCali    byte = 0x31 // pressure & temperature compensation calibration coefficients
	BMP388_RegPress   byte = 0x04 // start of pressure data registers
	BMP388_RegTemp    byte = 0x07 // start of temperature data registers
	BMP388_RegPwrCtrl byte = 0x1B // measurement mode & pressure/temperature sensor power register
	BMP388_RegOSR     byte = 0x1C // oversampling settings register
	BMP388_RegODR     byte = 0x1D //
	BMP388_RegCmd     byte = 0x7E // miscellaneous command register
	BMP388_RegStat    byte = 0x03 // sensor status register
	BMP388_RegErr     byte = 0x02 // error status register
	BMP388_RegIIR     byte = 0x1F
)

const (
	BMP388_ChipId    byte = 0x50 // correct response if reading from chip id register
	BMP388_PwrPress  byte = 0x01 // power on pressure sensor
	BMP388_PwrTemp   byte = 0x02 // power on temperature sensor
	BMP388_SoftReset byte = 0xB6 // command to reset all user configuration
	BMP388_DRDYPress byte = 0x20 // for checking if pressure data is ready
	BMP388_DRDYTemp  byte = 0x40 // for checking if pressure data is ready
)

// The difference between forced and normal mode is the bmp388 goes to sleep after taking a measurement in forced mode.
// Set it to forced if you intend to take measurements sporadically and want to save power. The driver will handle
// waking the sensor up when the sensor is in forced mode.
const (
	BMP388_Normal bmp388Mode = 0x30
	BMP388_Forced bmp388Mode = 0x16
	BMP388_Sleep  bmp388Mode = 0x00
)

// Increasing sampling rate increases precision but also the wait time for measurements. The datasheet has a table of
// suggested values for oversampling, output data rates, and iir filter coefficients by use case.
const (
	BMP388_Sampling1X bmp388Oversampling = iota
	BMP388_Sampling2X
	BMP388_Sampling4X
	BMP388_Sampling8X
	BMP388_Sampling16X
	BMP388_Sampling32X
)

// Output data rates in Hz. If increasing the sampling rates you need to decrease the output data rates, else the bmp388
// will freeze and Configure() will return a configuration error message. In that case keep decreasing the data rate
// until the bmp is happy
const (
	BMP388_Odr200 bmp388OutputDataRate = iota
	BMP388_Odr100
	BMP388_Odr50
	BMP388_Odr25
	BMP388_Odr12p5
	BMP388_Odr6p25
	BMP388_Odr3p1
	BMP388_Odr1p5
	BMP388_Odr0p78
	BMP388_Odr0p39
	BMP388_Odr0p2
	BMP388_Odr0p1
	BMP388_Odr0p05
	BMP388_Odr0p02
	BMP388_Odr0p01
	BMP388_Odr0p006
	BMP388_Odr0p003
	BMP388_Odr0p0015
)

// IIR filter coefficients, higher values means steadier measurements but slower reaction times
const (
	BMP388_Coeff0 bmp388FilterCoefficient = iota
	BMP388_Coeff1
	BMP388_Coeff3
	BMP388_Coeff7
	BMP388_Coeff15
	BMP388_Coeff31
	BMP388_Coeff63
	BMP388_Coeff127
)

// DS3231 registers
const Ds3231Address = 0x68

const (
	DS3231_REG_TIMEDATE = 0x00
	DS3231_REG_ALARMONE = 0x07
	DS3231_REG_ALARMTWO = 0x0B

	DS3231_REG_CONTROL = 0x0E
	DS3231_REG_STATUS  = 0x0F
	DS3231_REG_AGING   = 0x10

	DS3231_REG_TEMP = 0x11

	DS3231_REG_ALARMONE_SIZE = 4
	DS3231_REG_ALARMTWO_SIZE = 3

	// DS3231 Control Register Bits
	DS3231_A1IE  = 0
	DS3231_A2IE  = 1
	DS3231_INTCN = 2
	DS3231_RS1   = 3
	DS3231_RS2   = 4
	DS3231_CONV  = 5
	DS3231_BBSQW = 6
	DS3231_EOSC  = 7

	// DS3231 Status Register Bits
	DS3231_A1F     = 0
	DS3231_A2F     = 1
	DS3231_BSY     = 2
	DS3231_EN32KHZ = 3
	DS3231_OSF     = 7

	DS3231_AlarmFlag_Alarm1    = 0x01
	DS3231_AlarmFlag_Alarm2    = 0x02
	DS3231_AlarmFlag_AlarmBoth = 0x03

	DS3231_None          Ds3231Mode = 0
	DS3231_BatteryBackup Ds3231Mode = 1
	DS3231_Clock         Ds3231Mode = 2
	DS3231_AlarmOne      Ds3231Mode = 3
	DS3231_AlarmTwo      Ds3231Mode = 4
	DS3231_ModeAlarmBoth Ds3231Mode = 5
)

// INA250 registers
const Ina250Address = 0x40

const (
	INA250_REG_CONFIG     = 0x00
	INA250_REG_CURRENT    = 0x01
	INA250_REG_BUSVOLTAGE = 0x02
	INA250_REG_POWER      = 0x03
	INA250_REG_MASKENABLE = 0x06
	INA250_REG_ALERTLIMIT = 0x07
	INA250_REG_MANF_ID    = 0xFE
	INA250_REG_DIE_ID     = 0xFF
)

// Well-Known Values
const (
	INA250_MANF_ID        = 0x5449 // TI
	INA250_DEVICE_ID      = 0x2270 // 227h
	INA250_DEVICE_ID_MASK = 0xFFF0

	INA250_AVGMODE_1    = 0
	INA250_AVGMODE_4    = 1
	INA250_AVGMODE_16   = 2
	INA250_AVGMODE_64   = 3
	INA250_AVGMODE_128  = 4
	INA250_AVGMODE_256  = 5
	INA250_AVGMODE_512  = 6
	INA250_AVGMODE_1024 = 7

	INA250_CONVTIME_140USEC  = 0
	INA250_CONVTIME_204USEC  = 1
	INA250_CONVTIME_332USEC  = 2
	INA250_CONVTIME_588USEC  = 3
	INA250_CONVTIME_1100USEC = 4 // 1.1 ms
	INA250_CONVTIME_2116USEC = 5 // 2.1 ms
	INA250_CONVTIME_4156USEC = 6 // 4.2 ms
	INA250_CONVTIME_8244USEC = 7 // 8.2 ms

	INA250_MODE_CONTINUOUS = 0x4
	INA250_MODE_TRIGGERED  = 0x0
	INA250_MODE_VOLTAGE    = 0x2
	INA250_MODE_NO_VOLTAGE = 0x0
	INA250_MODE_CURRENT    = 0x1
	INA250_MODE_NO_CURRENT = 0x0
)

// LIS3DH registers
const (
	Lis3dhAddress0 = 0x18 // SA0 is low
	Lis3dhAddress1 = 0x19 // SA0 is high
)

const (
	LIS3DH_WHO_AM_I      = 0x0F
	LIS3DH_REG_STATUS1   = 0x07
	LIS3DH_REG_OUTADC1_L = 0x08
	LIS3DH_REG_OUTADC1_H = 0x09
	LIS3DH_REG_OUTADC2_L = 0x0A
	LIS3DH_REG_OUTADC2_H = 0x0B
	LIS3DH_REG_OUTADC3_L = 0x0C
	LIS3DH_REG_OUTADC3_H = 0x0D
	LIS3DH_REG_INTCOUNT  = 0x0E
	LIS3DH_REG_WHOAMI    = 0x0F
	LIS3DH_REG_TEMPCFG   = 0x1F
	LIS3DH_REG_CTRL1     = 0x20
	LIS3DH_REG_CTRL2     = 0x21
	LIS3DH_REG_CTRL3     = 0x22
	LIS3DH_REG_CTRL4     = 0x23
	LIS3DH_REG_CTRL5     = 0x24
	LIS3DH_REG_CTRL6     = 0x25
	LIS3DH_REG_REFERENCE = 0x26
	LIS3DH_REG_STATUS2   = 0x27
	LIS3DH_REG_OUT_X_L   = 0x28
	LIS3DH_REG_OUT_X_H   = 0x29
	LIS3DH_REG_OUT_Y_L   = 0x2A
	LIS3DH_REG_OUT_Y_H   = 0x2B
	LIS3DH_REG_OUT_Z_L   = 0x2C
	LIS3DH_REG_OUT_Z_H   = 0x2D
	LIS3DH_REG_FIFOCTRL  = 0x2E
	LIS3DH_REG_FIFOSRC   = 0x2F
	LIS3DH_REG_INT1CFG   = 0x30
	LIS3DH_REG_INT1SRC   = 0x31
	LIS3DH_REG_INT1THS   = 0x32
	LIS3DH_REG_INT1DUR   = 0x33
	LIS3DH_REG_CLICKCFG  = 0x38
	LIS3DH_REG_CLICKSRC  = 0x39
	LIS3DH_REG_CLICKTHS  = 0x3A
	LIS3DH_REG_TIMELIMIT = 0x3B
	LIS3DH_REG_TIMELATEN = 0x3C
	LIS3DH_REG_TIMEWINDO = 0x3D
	LIS3DH_REG_ACTTHS    = 0x3E
	LIS3DH_REG_ACTDUR    = 0x3F
)

type Lis3dhRange uint8

const (
	LIS3DH_RANGE_16_G Lis3dhRange = 3 // +/- 16g
	LIS3DH_RANGE_8_G              = 2 // +/- 8g
	LIS3DH_RANGE_4_G              = 1 // +/- 4g
	LIS3DH_RANGE_2_G              = 0 // +/- 2g (default value)
)

type Lis3dhDataRate uint8

// Data rate constants.
const (
	LIS3DH_DATARATE_400_HZ         Lis3dhDataRate = 7 //  400Hz
	LIS3DH_DATARATE_200_HZ                        = 6 //  200Hz
	LIS3DH_DATARATE_100_HZ                        = 5 //  100Hz
	LIS3DH_DATARATE_50_HZ                         = 4 //   50Hz
	LIS3DH_DATARATE_25_HZ                         = 3 //   25Hz
	LIS3DH_DATARATE_10_HZ                         = 2 // 10 Hz
	LIS3DH_DATARATE_1_HZ                          = 1 // 1 Hz
	LIS3DH_DATARATE_POWERDOWN                     = 0
	LIS3DH_DATARATE_LOWPOWER_1K6HZ                = 8
	LIS3DH_DATARATE_LOWPOWER_5KHZ                 = 9
)

// LPS22HB registers
const Lps22hbAddress = 0x5C

const (
	LPS22HB_WHO_AM_I_REG  = 0x0F
	LPS22HB_CTRL1_REG     = 0x10
	LPS22HB_CTRL2_REG     = 0x11
	LPS22HB_STATUS_REG    = 0x27
	LPS22HB_PRESS_OUT_REG = 0x28
	LPS22HB_TEMP_OUT_REG  = 0x2B
)

// LSM6DS3 registers
const Lsm6ds3Address = 0x6A

const (
	LSM6DS3_WHO_AM_I             = 0x0F
	LSM6DS3_STATUS               = 0x1E
	LSM6DS3_CTRL1_XL             = 0x10
	LSM6DS3_CTRL2_G              = 0x11
	LSM6DS3_CTRL3_C              = 0x12
	LSM6DS3_CTRL4_C              = 0x13
	LSM6DS3_CTRL5_C              = 0x14
	LSM6DS3_CTRL6_C              = 0x15
	LSM6DS3_CTRL7_G              = 0x16
	LSM6DS3_CTRL8_XL             = 0x17
	LSM6DS3_CTRL9_XL             = 0x18
	LSM6DS3_CTRL10_C             = 0x19
	LSM6DS3_OUTX_L_G             = 0x22
	LSM6DS3_OUTX_H_G             = 0x23
	LSM6DS3_OUTY_L_G             = 0x24
	LSM6DS3_OUTY_H_G             = 0x25
	LSM6DS3_OUTZ_L_G             = 0x26
	LSM6DS3_OUTZ_H_G             = 0x27
	LSM6DS3_OUTX_L_XL            = 0x28
	LSM6DS3_OUTX_H_XL            = 0x29
	LSM6DS3_OUTY_L_XL            = 0x2A
	LSM6DS3_OUTY_H_XL            = 0x2B
	LSM6DS3_OUTZ_L_XL            = 0x2C
	LSM6DS3_OUTZ_H_XL            = 0x2D
	LSM6DS3_OUT_TEMP_L           = 0x20
	LSM6DS3_OUT_TEMP_H           = 0x21
	LSM6DS3_BW_SCAL_ODR_DISABLED = 0x00
	LSM6DS3_BW_SCAL_ODR_ENABLED  = 0x80
	LSM6DS3_STEP_TIMESTAMP_L     = 0x49
	LSM6DS3_STEP_TIMESTAMP_H     = 0x4A
	LSM6DS3_STEP_COUNTER_L       = 0x4B
	LSM6DS3_STEP_COUNTER_H       = 0x4C
	LSM6DS3_STEP_COUNT_DELTA     = 0x15
	LSM6DS3_TAP_CFG              = 0x58
	LSM6DS3_INT1_CTRL            = 0x0D

	LSM6DS3_ACCEL_2G  lsm6ds3AccelRange = 0x00
	LSM6DS3_ACCEL_4G  lsm6ds3AccelRange = 0x08
	LSM6DS3_ACCEL_8G  lsm6ds3AccelRange = 0x0C
	LSM6DS3_ACCEL_16G lsm6ds3AccelRange = 0x04

	LSM6DS3_ACCEL_SR_OFF   lsm6ds3AccelSampleRate = 0x00
	LSM6DS3_ACCEL_SR_13    lsm6ds3AccelSampleRate = 0x10
	LSM6DS3_ACCEL_SR_26    lsm6ds3AccelSampleRate = 0x20
	LSM6DS3_ACCEL_SR_52    lsm6ds3AccelSampleRate = 0x30
	LSM6DS3_ACCEL_SR_104   lsm6ds3AccelSampleRate = 0x40
	LSM6DS3_ACCEL_SR_208   lsm6ds3AccelSampleRate = 0x50
	LSM6DS3_ACCEL_SR_416   lsm6ds3AccelSampleRate = 0x60
	LSM6DS3_ACCEL_SR_833   lsm6ds3AccelSampleRate = 0x70
	LSM6DS3_ACCEL_SR_1666  lsm6ds3AccelSampleRate = 0x80
	LSM6DS3_ACCEL_SR_3332  lsm6ds3AccelSampleRate = 0x90
	LSM6DS3_ACCEL_SR_6664  lsm6ds3AccelSampleRate = 0xA0
	LSM6DS3_ACCEL_SR_13330 lsm6ds3AccelSampleRate = 0xB0

	LSM6DS3_ACCEL_BW_50  lsm6ds3AccelBandwidth = 0x03
	LSM6DS3_ACCEL_BW_100 lsm6ds3AccelBandwidth = 0x02
	LSM6DS3_ACCEL_BW_200 lsm6ds3AccelBandwidth = 0x01
	LSM6DS3_ACCEL_BW_400 lsm6ds3AccelBandwidth = 0x00

	//GYRO_125DPS  GyroRange = 0x01
	LSM6DS3_GYRO_250DPS  lsm6ds3GyroRange = 0x00
	LSM6DS3_GYRO_500DPS  lsm6ds3GyroRange = 0x04
	LSM6DS3_GYRO_1000DPS lsm6ds3GyroRange = 0x08
	LSM6DS3_GYRO_2000DPS lsm6ds3GyroRange = 0x0C

	LSM6DS3_GYRO_SR_OFF  lsm6ds3GyroSampleRate = 0x00
	LSM6DS3_GYRO_SR_13   lsm6ds3GyroSampleRate = 0x10
	LSM6DS3_GYRO_SR_26   lsm6ds3GyroSampleRate = 0x20
	LSM6DS3_GYRO_SR_52   lsm6ds3GyroSampleRate = 0x30
	LSM6DS3_GYRO_SR_104  lsm6ds3GyroSampleRate = 0x40
	LSM6DS3_GYRO_SR_208  lsm6ds3GyroSampleRate = 0x50
	LSM6DS3_GYRO_SR_416  lsm6ds3GyroSampleRate = 0x60
	LSM6DS3_GYRO_SR_833  lsm6ds3GyroSampleRate = 0x70
	LSM6DS3_GYRO_SR_1666 lsm6ds3GyroSampleRate = 0x80
)

// MPU6050 registers
const Mpu6050Address = 0x68

const (
	// Self test registers
	MPU6050_SELF_TEST_X = 0x0D
	MPU6050_SELF_TEST_Y = 0x0E
	MPU6050_SELF_TEST_Z = 0x0F
	MPU6050_SELF_TEST_A = 0x10

	MPU6050_SMPLRT_DIV   = 0x19 // Sample rate divider
	MPU6050_CONFIG       = 0x1A // Configuration
	MPU6050_GYRO_CONFIG  = 0x1B // Gyroscope configuration
	MPU6050_ACCEL_CONFIG = 0x1C // Accelerometer configuration
	MPU6050_FIFO_EN      = 0x23 // FIFO enable

	// I2C pass-through configuration
	MPU6050_I2C_MST_CTRL   = 0x24
	MPU6050_I2C_SLV0_ADDR  = 0x25
	MPU6050_I2C_SLV0_REG   = 0x26
	MPU6050_I2C_SLV0_CTRL  = 0x27
	MPU6050_I2C_SLV1_ADDR  = 0x28
	MPU6050_I2C_SLV1_REG   = 0x29
	MPU6050_I2C_SLV1_CTRL  = 0x2A
	MPU6050_I2C_SLV2_ADDR  = 0x2B
	MPU6050_I2C_SLV2_REG   = 0x2C
	MPU6050_I2C_SLV2_CTRL  = 0x2D
	MPU6050_I2C_SLV3_ADDR  = 0x2E
	MPU6050_I2C_SLV3_REG   = 0x2F
	MPU6050_I2C_SLV3_CTRL  = 0x30
	MPU6050_I2C_SLV4_ADDR  = 0x31
	MPU6050_I2C_SLV4_REG   = 0x32
	MPU6050_I2C_SLV4_DO    = 0x33
	MPU6050_I2C_SLV4_CTRL  = 0x34
	MPU6050_I2C_SLV4_DI    = 0x35
	MPU6050_I2C_MST_STATUS = 0x36

	// Interrupt configuration
	MPU6050_INT_PIN_CFG = 0x37 // Interrupt pin/bypass enable configuration
	MPU6050_INT_ENABLE  = 0x38 // Interrupt enable
	MPU6050_INT_STATUS  = 0x3A // Interrupt status

	// Accelerometer measurements
	MPU6050_ACCEL_XOUT_H = 0x3B
	MPU6050_ACCEL_XOUT_L = 0x3C
	MPU6050_ACCEL_YOUT_H = 0x3D
	MPU6050_ACCEL_YOUT_L = 0x3E
	MPU6050_ACCEL_ZOUT_H = 0x3F
	MPU6050_ACCEL_ZOUT_L = 0x40

	// Temperature measurement
	MPU6050_TEMP_OUT_H = 0x41
	MPU6050_TEMP_OUT_L = 0x42

	// Gyroscope measurements
	MPU6050_GYRO_XOUT_H = 0x43
	MPU6050_GYRO_XOUT_L = 0x44
	MPU6050_GYRO_YOUT_H = 0x45
	MPU6050_GYRO_YOUT_L = 0x46
	MPU6050_GYRO_ZOUT_H = 0x47
	MPU6050_GYRO_ZOUT_L = 0x48

	// External sensor data
	MPU6050_EXT_SENS_DATA_00 = 0x49
	MPU6050_EXT_SENS_DATA_01 = 0x4A
	MPU6050_EXT_SENS_DATA_02 = 0x4B
	MPU6050_EXT_SENS_DATA_03 = 0x4C
	MPU6050_EXT_SENS_DATA_04 = 0x4D
	MPU6050_EXT_SENS_DATA_05 = 0x4E
	MPU6050_EXT_SENS_DATA_06 = 0x4F
	MPU6050_EXT_SENS_DATA_07 = 0x50
	MPU6050_EXT_SENS_DATA_08 = 0x51
	MPU6050_EXT_SENS_DATA_09 = 0x52
	MPU6050_EXT_SENS_DATA_10 = 0x53
	MPU6050_EXT_SENS_DATA_11 = 0x54
	MPU6050_EXT_SENS_DATA_12 = 0x55
	MPU6050_EXT_SENS_DATA_13 = 0x56
	MPU6050_EXT_SENS_DATA_14 = 0x57
	MPU6050_EXT_SENS_DATA_15 = 0x58
	MPU6050_EXT_SENS_DATA_16 = 0x59
	MPU6050_EXT_SENS_DATA_17 = 0x5A
	MPU6050_EXT_SENS_DATA_18 = 0x5B
	MPU6050_EXT_SENS_DATA_19 = 0x5C
	MPU6050_EXT_SENS_DATA_20 = 0x5D
	MPU6050_EXT_SENS_DATA_21 = 0x5E
	MPU6050_EXT_SENS_DATA_22 = 0x5F
	MPU6050_EXT_SENS_DATA_23 = 0x60

	// I2C peripheral data out
	MPU6050_I2C_PER0_DO      = 0x63
	MPU6050_I2C_PER1_DO      = 0x64
	MPU6050_I2C_PER2_DO      = 0x65
	MPU6050_I2C_PER3_DO      = 0x66
	MPU6050_I2C_MST_DELAY_CT = 0x67

	MPU6050_SIGNAL_PATH_RES = 0x68 // Signal path reset
	MPU6050_USER_CTRL       = 0x6A // User control
	MPU6050_PWR_MGMT_1      = 0x6B // Power Management 1
	MPU6050_PWR_MGMT_2      = 0x6C // Power Management 2
	MPU6050_FIFO_COUNTH     = 0x72 // FIFO count registers (high bits)
	MPU6050_FIFO_COUNTL     = 0x73 // FIFO count registers (low bits)
	MPU6050_FIFO_R_W        = 0x74 // FIFO read/write
	MPU6050_WHO_AM_I        = 0x75 // Who am I
)

// SHT3X registers
const (
	Sht3xAddressA = 0x44
	Sht3xAddressB = 0x45
)

const (
	// single shot, high repeatability
	SHT3X_MEASUREMENT_COMMAND_MSB = 0x24
	SHT3X_MEASUREMENT_COMMAND_LSB = 0x00
)

// VL53L1x registers
const Vl53l1xAddress = 0x29

const (
	VL53L1x_CHIP_ID                                                 = 0xEACC
	VL53L1x_SOFT_RESET                                              = 0x0000
	VL53L1x_OSC_MEASURED_FAST_OSC_FREQUENCY                         = 0x0006
	VL53L1x_VHV_CONFIG_TIMEOUT_MACROP_LOOP_BOUND                    = 0x0008
	VL53L1x_VHV_CONFIG_INIT                                         = 0x000B
	VL53L1x_ALGO_PART_TO_PART_RANGE_OFFSET_MM                       = 0x001E
	VL53L1x_MM_CONFIG_OUTER_OFFSET_MM                               = 0x0022
	VL53L1x_DSS_CONFIG_TARGET_TOTAL_RATE_MCPS                       = 0x0024
	VL53L1x_PAD_I2C_HV_EXTSUP_CONFIG                                = 0x002E
	VL53L1x_GPIO_TIO_HV_STATUS                                      = 0x0031
	VL53L1x_SIGMA_ESTIMATOR_EFFECTIVE_PULSE_WIDTH_NS                = 0x0036
	VL53L1x_SIGMA_ESTIMATOR_EFFECTIVE_AMBIENT_WIDTH_NS              = 0x0037
	VL53L1x_ALGO_CROSSTALK_COMPENSATION_VALID_HEIGHT_MM             = 0x0039
	VL53L1x_ALGO_RANGE_MIN_CLIP                                     = 0x003F
	VL53L1x_ALGO_CONSISTENCY_CHECK_TOLERANCE                        = 0x0040
	VL53L1x_CAL_CONFIG_VCSEL_START                                  = 0x0047
	VL53L1x_PHASECAL_CONFIG_TIMEOUT_MACROP                          = 0x004B
	VL53L1x_PHASECAL_CONFIG_OVERRIDE                                = 0x004D
	VL53L1x_DSS_CONFIG_ROI_MODE_CONTROL                             = 0x004F
	VL53L1x_SYSTEM_THRESH_RATE_HIGH                                 = 0x0050
	VL53L1x_SYSTEM_THRESH_RATE_LOW                                  = 0x0052
	VL53L1x_DSS_CONFIG_MANUAL_EFFECTIVE_SPADS_SELECT                = 0x0054
	VL53L1x_DSS_CONFIG_APERTURE_ATTENUATION                         = 0x0057
	VL53L1x_MM_CONFIG_TIMEOUT_MACROP_A                              = 0x005A
	VL53L1x_MM_CONFIG_TIMEOUT_MACROP_B                              = 0x005C
	VL53L1x_RANGE_CONFIG_TIMEOUT_MACROP_A                           = 0x005E
	VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A                             = 0x0060
	VL53L1x_RANGE_CONFIG_TIMEOUT_MACROP_B                           = 0x0061
	VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_B                             = 0x0063
	VL53L1x_RANGE_CONFIG_SIGMA_THRESH                               = 0x0064
	VL53L1x_RANGE_CONFIG_MIN_COUNT_RATE_RTN_LIMIT_MCPS              = 0x0066
	VL53L1x_RANGE_CONFIG_VALID_PHASE_HIGH                           = 0x0069
	VL53L1x_SYSTEM_INTERMEASUREMENT_PERIOD                          = 0x006C
	VL53L1x_SYSTEM_GROUPED_PARAMETER_HOLD_0                         = 0x0071
	VL53L1x_SYSTEM_SEED_CONFIG                                      = 0x0077
	VL53L1x_SD_CONFIG_WOI_SD0                                       = 0x0078
	VL53L1x_SD_CONFIG_WOI_SD1                                       = 0x0079
	VL53L1x_SD_CONFIG_INITIAL_PHASE_SD0                             = 0x007A
	VL53L1x_SD_CONFIG_INITIAL_PHASE_SD1                             = 0x007B
	VL53L1x_SYSTEM_GROUPED_PARAMETER_HOLD_1                         = 0x007C
	VL53L1x_SD_CONFIG_QUANTIFIER                                    = 0x007E
	VL53L1x_SYSTEM_SEQUENCE_CONFIG                                  = 0x0081
	VL53L1x_SYSTEM_GROUPED_PARAMETER_HOLD                           = 0x0082
	VL53L1x_SYSTEM_INTERRUPT_CLEAR                                  = 0x0086
	VL53L1x_SYSTEM_MODE_START                                       = 0x0087
	VL53L1x_RESULT_RANGE_STATUS                                     = 0x0089
	VL53L1x_PHASECAL_RESULT_VCSEL_START                             = 0x00D8
	VL53L1x_RESULT_OSC_CALIBRATE_VAL                                = 0x00DE
	VL53L1x_FIRMWARE_SYSTEM_STATUS                                  = 0x00E5
	VL53L1x_WHO_AM_I                                                = 0x010F
	VL53L1x_SHADOW_RESULT_FINAL_CROSSTALK_CORRECTED_RANGE_MM_SD0_HI = 0x0FBE

	VL53L1x_TIMING_GUARD = 4528
	VL53L1x_TARGETRATE   = 0x0A00
)

const (
	VL53L1x_SHORT vl53l1xDistanceMode = iota
	VL53L1x_MEDIUM
	VL53L1x_LONG
)

const (
	VL53L1x_RangeValid vl53l1xRangeStatus = iota
	VL53L1x_SigmaFail
	VL53L1x_SignalFail
	VL53L1x_RangeValidMinRangeClipped
	VL53L1x_OutOfBoundsFail
	VL53L1x_HardwareFail
	VL53L1x_RangeValidNoWrapCheckFail
	VL53L1x_WrapTargetFail
	VL53L1x_ProcessingFail
	VL53L1x_XtalkSignalFail
	VL53L1x_SynchronizationInt
	VL53L1x_MergedPulse
	VL53L1x_TargetPresentLackOfSignal
	VL53L1x_MinRangeFail
	VL53L1x_RangeInvalid

	VL53L1x_None vl53l1xRangeStatus = 255
)
