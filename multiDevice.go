package multi

import (
	"machine"

	"tinygo.org/x/drivers"
)

type device interface {
	configure()
	connect()
}

func initializeDeviceMap(bus drivers.I2C, addr uint16) map[string]interface{} {

	deviceMap := make(map[string]interface{})

	deviceMap["adxl345"] = newAdx1345(bus, addr)
	deviceMap["amg88xx"] = newAmg88xx(bus, addr)
	deviceMap["at24cx"] = newAt24cx(bus, addr)
	deviceMap["bh1750"] = newBh1750(bus, addr)
	deviceMap["blinkm"] = newBlinkm(bus, addr)
	deviceMap["bme280"] = newBme280(bus, addr)
	deviceMap["bmp280"] = newBmp280(bus, addr)
	deviceMap["bmp388"] = newBmp388(bus, addr)
	deviceMap["ds3231"] = newDs3231(bus, addr)
	deviceMap["ina260"] = newIna260(bus, addr)
	deviceMap["lis3dh"] = newLis3dh(bus, addr)
	deviceMap["lps22hb"] = newLps22hb(bus, addr)
	deviceMap["lsm6ds3"] = newLsm6ds3(bus, addr)
	deviceMap["mpu6050"] = newMpu6050(bus, addr)
	deviceMap["sht3x"] = newSht3x(bus, addr)
	deviceMap["vl53l1x"] = newVl53l1x(bus, addr)

	return deviceMap
}

func NewDevice(mach *machine.I2C, deviceName string, addr uint16) {
	deviceMap := initializeDeviceMap(mach, addr)

	dev := deviceMap[deviceName]

	dev.(device).configure()
}
