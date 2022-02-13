package multi

import (
	"tinygo.org/x/drivers"
)

// TODO: set up device registers to that
// these device creation functions can
// handle default I2C addresses-

// TODO: REMOVE & Veml6070

func newAdx1345(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Adxl345{
			bus: bus,
			powerCtl: powerCtl{
				measure: 1,
			},
			dataFormat: dataFormat{
				// sensorRange: RANGE_2G,
			},
			bwRate: bwRate{
				lowPower: 1,
				// rate:     RATE_100HZ,
			},
			Address: addr,
		}
	} else {
		return Adxl345{
			bus: bus,
			powerCtl: powerCtl{
				measure: 1,
			},
			dataFormat: dataFormat{
				// sensorRange: RANGE_2G,
			},
			bwRate: bwRate{
				lowPower: 1,
				// rate:     RATE_100HZ,
			},
			// Address: AddressLow,
		}
	}
}

func newAht20(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Aht20{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Aht20{
			bus: bus,
			// Address: Address,
		}
	}
}

func newAmg88xx(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Aht20{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Aht20{
			bus: bus,
			// Address: AddressHigh,
		}
	}
}

func newApds9960(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Apds9960{
			bus:     bus,
			Address: uint8(addr),
			// mode: MODE_NONE,
		}
	} else {
		return Apds9960{
			bus: bus,
			// Address: ADPS9960_ADDRESS,
			// mode: MODE_NONE,
		}
	}
}

func newAt24cx(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return At24cx{
			bus:     bus,
			Address: addr,
		}
	} else {
		return At24cx{
			bus: bus,
			// Address: Address,
		}
	}
}

func newBh1750(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Bh1750{
			bus:     bus,
			Address: addr,
			// mode:    CONTINUOUS_HIGH_RES_MODE,
		}
	} else {
		return Bh1750{
			bus: bus,
			// Address: Address,
			// mode:    CONTINUOUS_HIGH_RES_MODE,
		}
	}
}

func newBlinkm(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Blinkm{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Blinkm{
			bus: bus,
			// Address: AddressHigh,
		}
	}
}

func newBme280(bus drivers.I2C, addr uint16) interface{} {

	if addr == 0 {
		return Bme280{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Bme280{
			bus: bus,
			// Address: Address,
		}
	}
}

func newBmp180(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Bmp180{
			bus:     bus,
			Address: addr,
			// mode:    ULTRAHIGHRESOLUTION,
		}
	} else {
		return Bmp180{
			bus: bus,
			// Address: Address,
			// mode:    ULTRAHIGHRESOLUTION,
		}
	}
}

func newBmp280(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Bmp280{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Bmp280{
			bus: bus,
			// Address: Address,
		}
	}
}

func newBmp388(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Bmp388{
			bus:     bus,
			Address: uint8(addr),
		}
	} else {
		return Bmp388{
			bus: bus,
			// Address: Address,
		}
	}
}

func newDs3231(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Ds3231{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Ds3231{
			bus: bus,
			// Address: Address,
		}
	}
}

func newHts221(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Hts221{
			bus:     bus,
			Address: uint8(addr),
		}
	} else {
		return Hts221{
			bus: bus,
			// Address: Address,
		}
	}
}

func newIna260(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Ina260{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Ina260{
			bus: bus,
			// Address: Address,
		}
	}
}

func newLis2mdl(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Lis2mdl{
			bus:     bus,
			Address: uint8(addr),
		}
	} else {
		return Lis2mdl{
			bus: bus,
			// Address: ADDRESS,
		}
	}
}

func newLis3dh(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Lis3dh{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Lis2mdl{
			bus: bus,
			// Address: Address0,
		}
	}
}

func newLps22hb(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Lis3dh{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Lis2mdl{
			bus: bus,
			// Address: LPS22HB_ADDRESS,
		}
	}
}

func newLsm6ds3(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Lsm6ds3{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Lsm6ds3{
			bus: bus,
			// Address: Address,
		}
	}
}

func newMag3110(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Mag3110{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Mag3110{
			bus: bus,
			// Address: Address,
		}
	}
}

func newMma8653(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Mma8653{
			bus:     bus,
			Address: addr,
			// sensitivity: Sensitivity2G,
		}
	} else {
		return Mma8653{
			bus: bus,
			// Address: Address,
			// sensitivity: Sensitivity2G,
		}
	}
}

func newMpu6050(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Mpu6050{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Mpu6050{
			bus: bus,
			// Address: Address,
		}
	}
}

func newPcf8563(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Pcf8563{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Pcf8563{
			bus: bus,
			// Address: PCF8563_ADDR,
		}
	}
}

func newSht3x(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Sht3x{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Sht3x{
			bus: bus,
			// Address: AddressA,
		}
	}
}

// Will need to look into multi-address for Tmp102
// creation more (tinygo/drivers/tmp120 uses the
// Config method for setting up multi-addresses)
func newTmp102(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	return nil
}

func newVl53l1x(bus drivers.I2C, addr uint16) interface{} {
	//TODO: Add creation of device to be returned

	if addr == 0 {
		return Vl53l1x{
			bus:     bus,
			Address: addr,
			// mode: LONG,
			timeout: 500,
		}
	} else {
		return Vl53l1x{
			bus: bus,
			// Address: Address,
			// mode: LONG,
			timeout: 500,
		}
	}
}
