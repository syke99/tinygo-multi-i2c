package v2

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

// BH1750 registers
const Bh1750Address_1 = 0x23
const Bh1750Address_2 = 0x23

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
const BlinkmAddress_DEFAULT = 0x09
const BlinkmAddress_GENERAL = 0x00

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
const Bme280Address_1 = 0x76
const Bme280Address_2 = 0x77

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
const Lps22hbAddress_1 = 0x5C
const Lps22hbAddress_2 = 0x5D

const (
	LPS22HB_WHO_AM_I_REG  = 0x0F
	LPS22HB_CTRL1_REG     = 0x10
	LPS22HB_CTRL2_REG     = 0x11
	LPS22HB_STATUS_REG    = 0x27
	LPS22HB_PRESS_OUT_REG = 0x28
	LPS22HB_TEMP_OUT_REG  = 0x2B
)

// MPU6050 registers
const Mpu6050Address_1 = 0x68
const Mpu6050Address_2 = 0x69

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
