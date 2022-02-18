package v2

import (
	"errors"
)

// The BMP280 Temperature, Humidity, and Berometric Pressure sensor by Bosch
// Sensortech requires some settings to be passed in as an array with
// a length of 5, of uint values to set Standby, Filter, Temperature, Pressure,
// and Mode
func NewDevice(bus I2C, deviceName string, devNumber int, bmp280Settings [5]uint) (Devices, []error) {

	adxl345Addresses := []uint16{Adx1345AddressHigh, Adx1345AddressLow}
	amg88xxAddresses := []uint16{Amg88xxAddressHigh, Amg88xxAddressLow}
	bh1650Addresses := []uint16{Bh1750Address_1, Bh1750Address_2}
	blinkmAddresses := []uint16{BlinkmAddress_DEFAULT, BlinkmAddress_GENERAL}
	bme280Addresses := []uint16{Bme280Address_1, Bme280Address_2}
	lis3dhAddresses := []uint16{Lis3dhAddress0, Lis3dhAddress1}
	lps22hbAddresses := []uint16{Lps22hbAddress_1, Lps22hbAddress_2}
	mpu6050Addresses := []uint16{Mpu6050Address_1, Mpu6050Address_2}

	var dvc Devices

	var newDeviceError []error

	switch deviceName {
	case "adxl345":

		i := 0

		if devNumber <= len(adxl345Addresses) {
			for i < devNumber {
				dvc.addAdxl1345(newAdx1345(bus, adxl345Addresses[i]))

				dvc.adxl345[i].configure()
			}
			return dvc, newDeviceError
		} else {
			createErr := errors.New("number of devices to create greater than available addresses for device")
			return dvc, append(newDeviceError, createErr)
		}
	case "amg88xx":

		i := 0

		if devNumber <= len(amg88xxAddresses) {
			for i < devNumber {
				dvc.addAmg88xx(newAmg88xx(bus, amg88xxAddresses[i]))

				dvc.amg88xx[i].configure()
			}
			return dvc, newDeviceError
		} else {
			createErr := errors.New("number of devices to create greater than available addresses for device")
			return dvc, append(newDeviceError, createErr)
		}
	case "bh1750":

		i := 0

		if devNumber <= len(bh1650Addresses) {
			for i < devNumber {
				dvc.addBh1750(newBh1750(bus, bh1650Addresses[i]))

				dvc.bh1750[i].configure()
			}
			return dvc, newDeviceError
		} else {
			createErr := errors.New("number of devices to create greater than available addresses for device")
			return dvc, append(newDeviceError, createErr)
		}
	case "blinkm":

		i := 0

		if devNumber <= len(blinkmAddresses) {
			for i < devNumber {
				dvc.addBlinkm(newBlinkm(bus, blinkmAddresses[i]))

				dvc.blinkm[i].configure()
			}
			return dvc, newDeviceError
		} else {
			createErr := errors.New("number of devices to create greater than available addresses for device")
			return dvc, append(newDeviceError, createErr)
		}
	case "bme280":

		i := 0

		if devNumber <= len(bme280Addresses) {
			for i < devNumber {
				dvc.addBme280(newBme280(bus, bme280Addresses[i]))

				dvc.bme280[i].configure()
			}

			err := dvc.bme280[i].connected()
			if !err {
				newDeviceError = append(newDeviceError, errors.New("device configured but unable to connect"))
			}

			return dvc, newDeviceError
		} else {
			createErr := errors.New("number of devices to create greater than available addresses for device")
			return dvc, append(newDeviceError, createErr)
		}
	case "lis3dh":

		i := 0

		if devNumber <= len(lis3dhAddresses) {
			for i < devNumber {
				dvc.addLis3dh(newLis3dh(bus, lis3dhAddresses[i]))

				dvc.lis3dh[i].configure()
			}

			err := dvc.lis3dh[i].connected()
			if !err {
				newDeviceError = append(newDeviceError, errors.New("device configured but unable to connect"))
			}

			return dvc, newDeviceError
		} else {
			createErr := errors.New("number of devices to create greater than available addresses for device")
			return dvc, append(newDeviceError, createErr)
		}
	case "lps22hb":

		i := 0

		if devNumber <= len(lps22hbAddresses) {
			for i < devNumber {
				dvc.addLps22hb(newLps22hb(bus, lps22hbAddresses[i]))

				dvc.lps22hb[i].configure()
			}

			err := dvc.lps22hb[i].connected()
			if !err {
				newDeviceError = append(newDeviceError, errors.New("device configured but unable to connect"))
			}

			return dvc, newDeviceError
		} else {
			createErr := errors.New("number of devices to create greater than available addresses for device")
			return dvc, append(newDeviceError, createErr)
		}
	case "mpu6050":

		i := 0

		if devNumber <= len(mpu6050Addresses) {
			for i < devNumber {
				dvc.addMpu6050(newMpu6050(bus, mpu6050Addresses[i]))

				dvc.mpu6050[i].configure()
			}

			err := dvc.mpu6050[i].connected()
			if !err {
				newDeviceError = append(newDeviceError, errors.New("device configured but unable to connect"))
			}

			return dvc, newDeviceError
		} else {
			createErr := errors.New("number of devices to create greater than available addresses for device")
			return dvc, append(newDeviceError, createErr)
		}
	default:
		return dvc, newDeviceError
	}

}
