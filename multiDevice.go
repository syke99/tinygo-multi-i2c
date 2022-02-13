package multi

import (
	"fmt"
	"machine"

	"tinygo.org/x/drivers"
)

func initializeDeviceMap(bus drivers.I2C, addr uint16) map[string]interface{} {

	deviceMap := make(map[string]interface{})

	deviceMap["adxl345"] = newAdx1345(bus, addr)
	deviceMap["aht20"] = newAht20(bus, addr)
	deviceMap["amg88xx"] = newAmg88xx(bus, addr)
	deviceMap["apds9960"] = newApds9960(bus, addr)
	deviceMap["at24cx"] = newAt24cx(bus, addr)
	deviceMap["bh1750"] = newBh1750(bus, addr)
	deviceMap["blinkm"] = newBlinkm(bus, addr)
	deviceMap["bme280"] = newBme280(bus, addr)
	deviceMap["bmp180"] = newBmp180(bus, addr)
	deviceMap["bmp280"] = newBmp280(bus, addr)
	deviceMap["bmp388"] = newBmp388(bus, addr)
	deviceMap["ds3231"] = newDs3231(bus, addr)
	deviceMap["hts221"] = newHts221(bus, addr)
	deviceMap["ina260"] = newIna260(bus, addr)
	deviceMap["lis2mdl"] = newLis2mdl(bus, addr)
	deviceMap["lis3dh"] = newLis3dh(bus, addr)
	deviceMap["lps22hb"] = newLps22hb(bus, addr)
	deviceMap["lsm6ds3"] = newLsm6ds3(bus, addr)
	deviceMap["mag3110"] = newMag3110(bus, addr)
	deviceMap["mma8653"] = newMma8653(bus, addr)
	deviceMap["mpu6050"] = newMpu6050(bus, addr)
	deviceMap["pcf8563"] = newPcf8563(bus, addr)
	deviceMap["sht3x"] = newSht3x(bus, addr)
	deviceMap["tmp102"] = newTmp102(bus, addr)
	deviceMap["vl53l1xDevice"] = newVl53l1x(bus, addr)

	return deviceMap
}

func NewDevice(mach *machine.I2C, deviceName string, addr uint16) {
	// machine.I2C0.Configure(machine.I2CConfig{})
	// sensor := bmp180.New(machine.I2C0)
	// sensor.Configure()

	// sensor.Address = addr

	deviceMap := initializeDeviceMap(mach, addr)

	dev := deviceMap[deviceName]

	fmt.Println(dev)
}
