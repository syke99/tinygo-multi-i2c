package multi

// TODO: flesh out connect methods for each device

func (d Bme280) connected() bool {
	data := []byte{0}
	d.bus.ReadRegister(uint8(d.Address), BME280_WHO_AM_I, data)
	return data[0] == BME280_CHIP_ID
}

func (d Bmp280) connected() bool {
	data := make([]byte, 1)
	d.bus.ReadRegister(uint8(d.Address), BMP280_REG_ID, data)
	return data[0] == BMP280_CHIP_ID
}

func (d Bmp388) connected() bool {
	data, err := d.readRegister(BMP388_RegChipId, 1)
	return err == nil && data[0] == BMP388_ChipId // returns true if i2c comm was good and response equals 0x50
}

func (i Ina260) connected() error {
	// Ina260 may need to be removed?
	return nil
}

func (d Lis3dh) connected() bool {
	data := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), LIS3DH_WHO_AM_I, data)
	if err != nil {
		return false
	}
	return data[0] == 0x33
}

func (d Lps22hb) connected() bool {
	data := []byte{0}
	d.bus.ReadRegister(d.Address, LPS22HB_WHO_AM_I_REG, data)
	return data[0] == 0xB1
}

func (d Mpu6050) connected() bool {
	data := []byte{0}
	d.bus.ReadRegister(uint8(d.Address), MPU6050_WHO_AM_I, data)
	return data[0] == 0x68
}
