package multi

import (
	"tinygo.org/x/drivers"
)

// TODO: set up device registers to that
// these device creation functions can
// handle default I2C addresses

func newAdx1345(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Adxl345{
			bus: bus,
			powerCtl: powerCtl{
				measure: 1,
			},
			dataFormat: adxl345DataFormat{
				// sensorRange: RANGE_2G,
			},
			bwRate: adxl345BwRate{
				lowPower: 1,
				rate:     ADX1345_RATE_100HZ,
			},
			Address: addr,
		}
	} else {
		return Adxl345{
			bus: bus,
			powerCtl: powerCtl{
				measure: 1,
			},
			dataFormat: adxl345DataFormat{
				// sensorRange: RANGE_2G,
			},
			bwRate: adxl345BwRate{
				lowPower: 1,
				rate:     ADX1345_RATE_100HZ,
			},
			Address: Adx1345AddressLow,
		}
	}
}

func newAmg88xx(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Amg88xx{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Amg88xx{
			bus:     bus,
			Address: Amg88xxAddressHigh,
		}
	}
}

func newAt24cx(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return At24cx{
			bus:     bus,
			Address: addr,
		}
	} else {
		return At24cx{
			bus:     bus,
			Address: At24cxAddress,
		}
	}
}

func newBh1750(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Bh1750{
			bus:     bus,
			Address: addr,
			mode:    BH1750_CONTINUOUS_HIGH_RES_MODE,
		}
	} else {
		return Bh1750{
			bus:     bus,
			Address: Bh1750Address,
			mode:    BH1750_CONTINUOUS_HIGH_RES_MODE,
		}
	}
}

func newBlinkm(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Blinkm{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Blinkm{
			bus:     bus,
			Address: BlinkmAddress,
		}
	}
}

func newBme280(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Bme280{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Bme280{
			bus:     bus,
			Address: Bme280Address,
		}
	}
}

func newBmp280(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Bmp280{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Bmp280{
			bus:     bus,
			Address: Bmp280Address,
		}
	}
}

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

func newDs3231(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Ds3231{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Ds3231{
			bus:     bus,
			Address: Ds3231Address,
		}
	}
}

func newIna260(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Ina260{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Ina260{
			bus:     bus,
			Address: Ina250Address,
		}
	}
}

func newLis3dh(bus drivers.I2C, addr uint16) interface{} {
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

func newLps22hb(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Lps22hb{
			bus:     bus,
			Address: uint8(addr),
		}
	} else {
		return Lps22hb{
			bus:     bus,
			Address: Lps22hbAddress,
		}
	}
}

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

func newSht3x(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Sht3x{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Sht3x{
			bus:     bus,
			Address: Sht3xAddressA,
		}
	}
}

func newVl53l1x(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Vl53l1x{
			bus:     bus,
			Address: addr,
			mode:    VL53L1x_LONG,
			timeout: 500,
		}
	} else {
		return Vl53l1x{
			bus:     bus,
			Address: Vl53l1xAddress,
			mode:    VL53L1x_LONG,
			timeout: 500,
		}
	}
}
