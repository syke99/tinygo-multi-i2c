package multi

import (
	"errors"
)

// The BMP280 Temperature, Humidity, and Berometric Pressure sensor by Bosch
// Sensortech requires some settings to be passed in as an array with
// a length of 5, of uint values to set Standby, Filter, Temperature, Pressure,
// and Mode
func NewDevice(bus I2C, deviceName string, addr uint16, bmp280Settings [5]uint) (Devices, error) {

	var dvc Devices

	var newDeviceError error

	// This new switch allows for avoiding unnecessary polymorphism
	// like was implemented in v0.1.0 the polymorphism required
	// switching on an interface type assertion that could cause
	// package to panic due to possibly removing the ability for more type
	// asserting that was required for testing the connection.
	switch deviceName {
	case "adxl345":
		dvc.addAdxl1345(newAdx1345(bus, addr))

		dvc.adxl345.configure()

		return dvc, newDeviceError
	case "amg88xx":
		dvc.addAmg88xx(newAmg88xx(bus, addr))

		dvc.amg88xx.configure()

		return dvc, newDeviceError
	case "bh1750":
		dvc.addBh1750(newBh1750(bus, addr))

		dvc.bh1750.configure()

		return dvc, newDeviceError
	case "blinkm":
		dvc.addBlinkm(newBlinkm(bus, addr))

		dvc.blinkm.configure()

		return dvc, newDeviceError
	case "bme280":
		dvc.addBme280(newBme280(bus, addr))

		dvc.bme280.configure()

		err := dvc.bme280.connected()
		if !err {
			newDeviceError = errors.New("device configured but unable to connect")
		}

		return dvc, newDeviceError
	case "bmp280":
		dvc.addBmp280(newBmp280(bus, addr))

		dvc.bmp280.configure(bmp280Settings[0], bmp280Settings[1], bmp280Settings[2], bmp280Settings[3], bmp280Settings[4])

		err := dvc.bmp280.connected()
		if !err {
			newDeviceError = errors.New("device configured but unable to connect")
		}

		return dvc, newDeviceError
	case "lis3dh":
		dvc.addLis3dh(newLis3dh(bus, addr))

		dvc.lis3dh.configure()

		err := dvc.lis3dh.connected()
		if !err {
			newDeviceError = errors.New("device configured but unable to connect")
		}

		return dvc, newDeviceError
	case "lps22hb":
		dvc.addLps22hb(newLps22hb(bus, addr))

		dvc.lps22hb.configure()

		err := dvc.lps22hb.connected()
		if !err {
			newDeviceError = errors.New("device configured but unable to connect")
		}

		return dvc, newDeviceError
	case "mpu6050":
		dvc.addMpu6050(newMpu6050(bus, addr))

		dvc.mpu6050.configure()

		err := dvc.mpu6050.connected()
		if !err {
			newDeviceError = errors.New("device configured but unable to connect")
		}

		return dvc, newDeviceError
	default:
		return dvc, newDeviceError
	}

}
