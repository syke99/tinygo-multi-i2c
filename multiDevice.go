package multi

import (
	"machine"

	"tinygo.org/x/drivers/bmp180"
	// "tinygo.org/x/drivers"
)

func InitializeDeviceMap() map[string]interface{} {

	deviceMap := make(map[string]interface{})

	deviceMap["adt740"] = new(Adt7410)
	deviceMap["adxl345"] = new(Adxl345)
	deviceMap["aht20"] = new(Aht20)
	deviceMap["amg88xx"] = new(Amg88xx)
	deviceMap["apds9960"] = new(Apds9960)
	deviceMap["at24cx"] = new(At24cx)
	deviceMap["bh1750"] = new(Bh1750)
	deviceMap["blinkm"] = new(Blinkm)
	deviceMap["bme280"] = new(Bme280)
	deviceMap["bmp180"] = new(Bmp180)
	deviceMap["bmp280"] = new(Bmp280)
	deviceMap["bmp388"] = new(Bmp388)
	deviceMap["ds1407"] = new(Ds1407)
	deviceMap["ds3231"] = new(Ds3231)
	deviceMap["ft6336"] = new(Ft6336)
	deviceMap["hd44780i2c"] = new(Hd44780i2c)
	deviceMap["hts221"] = new(Hts221)
	deviceMap["ina260"] = new(Ina260)
	deviceMap["lis2mdl"] = new(Lis2mdl)
	deviceMap["lis3dh"] = new(Lis3dh)
	deviceMap["lps22hb"] = new(Lps22hb)
	deviceMap["lsm6ds3"] = new(Lsm6ds3)
	deviceMap["lsm6dsox"] = new(Lsm6dsox)
	deviceMap["lsm9ds1"] = new(Lsm9ds1)
	deviceMap["lsm303agr"] = new(Lsm303agr)
	deviceMap["mag3110"] = new(Mag3110)
	deviceMap["mma8653"] = new(Mma8653)
	deviceMap["mpu6050"] = new(Mpu6050)
	deviceMap["pcf8563"] = new(Pcf8563)
	deviceMap["sht3x"] = new(Sht3x)
	deviceMap["tmp102"] = new(Tmp102)
	deviceMap["veml6070"] = new(Veml6070)
	deviceMap["vl53l1xDevice"] = new(Vl53l1xDevice)

	return deviceMap
}

func NewDevice(mach *machine.I2C, addr uint16, deviceMap map[string]interface{}) {
	machine.I2C0.Configure(machine.I2CConfig{})
	sensor := bmp180.New(machine.I2C0)
	sensor.Configure()

	sensor.Address = addr
}
