package multi

import (
	"errors"
	"fmt"
	"machine"

	"tinygo.org/x/drivers"
)

type device interface {
	configure(uint, uint, uint, uint) error
	connected() bool
}

func initializeDeviceMap(bus drivers.I2C, addr uint16) map[string]interface{} {

	deviceMap := make(map[string]interface{})

	deviceMap["adxl345"] = newAdx1345(bus, addr)
	deviceMap["amg88xx"] = newAmg88xx(bus, addr)
	deviceMap["bh1750"] = newBh1750(bus, addr)
	deviceMap["blinkm"] = newBlinkm(bus, addr)
	deviceMap["bme280"] = newBme280(bus, addr)
	deviceMap["bmp280"] = newBmp280(bus, addr)
	deviceMap["ds3231"] = newDs3231(bus, addr)
	deviceMap["ina260"] = newIna260(bus, addr)
	deviceMap["lis3dh"] = newLis3dh(bus, addr)
	deviceMap["lps22hb"] = newLps22hb(bus, addr)
	deviceMap["mpu6050"] = newMpu6050(bus, addr)

	return deviceMap
}

func NewDevice(mach *machine.I2C, deviceName string, addr uint16, bmp280Settings []uint) (Devices, error) {

	var newDeviceError error

	deviceMap := initializeDeviceMap(mach, addr)

	dev := deviceMap[deviceName]

	switch dev.(type) {
	case Bme280:
		newDeviceError = dev.(device).configure(bmp280Settings[0], bmp280Settings[1], bmp280Settings[2], bmp280Settings[3])
	default:
		newDeviceError = dev.(device).configure(0, 0, 0, 0)
	}

	connectedBool := false

	switch dev.(type) {
	case Bme280, Bmp280, Lis3dh, Lps22hb, Mpu6050:
		connectedBool = dev.(device).connected()
		if !connectedBool {
			newDeviceError = errors.New("Device configured but unable to connect.")
		}
	}

	var dvc Devices

	switch dev.(type) {
	case Adxl345:
		value, ok := dev.(Adxl345)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Adxl345)))
		} else {
			dvc.addAdxl1345(value)
		}
	case Amg88xx:
		value, ok := dev.(Amg88xx)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Amg88xx)))
		} else {
			dvc.addAmg88xx(value)
		}
	case Bh1750:
		value, ok := dev.(Bh1750)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Bh1750)))
		} else {
			dvc.addBh1750(value)
		}
	case Blinkm:
		value, ok := dev.(Blinkm)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Blinkm)))
		} else {
			dvc.addBlinkm(value)
		}
	case Bme280:
		value, ok := dev.(Bme280)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Bme280)))
		} else {
			dvc.addBme280(value)
		}
	case Bmp280:
		value, ok := dev.(Bmp280)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Bmp280)))
		} else {
			dvc.addBmp280(value)
		}
	case Ds3231:
		value, ok := dev.(Ds3231)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Ds3231)))
		} else {
			dvc.addDs3231(value)
		}
	case Ina260:
		value, ok := dev.(Ina260)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Ina260)))
		} else {
			dvc.addIna260(value)
		}
	case Lis3dh:
		value, ok := dev.(Lis3dh)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Lis3dh)))
		} else {
			dvc.addLis3dh(value)
		}
	case Lps22hb:
		value, ok := dev.(Lps22hb)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Lps22hb)))
		} else {
			dvc.addLps22hb(value)
		}
	case Mpu6050:
		value, ok := dev.(Mpu6050)

		if !ok {
			newDeviceError = errors.New(fmt.Sprintf("Cannot set created device to type: %T. Device creation failed.", dev.(Mpu6050)))
		} else {
			dvc.addMpu6050(value)
		}
	default:
		newDeviceError = errors.New("Device creation failure.")
	}

	return dvc, newDeviceError
}
