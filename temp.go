// TODO: port over device registers for adding to new device

package multi

type Devices struct {
	adt740        Adt7410
	adxl345       Adxl345
	aht20         Aht20
	amg88xx       Amg88xx
	apds9960      Apds9960
	at24cx        At24cx
	bh1750        Bh1750
	blinkm        Blinkm
	bme280        Bme280
	bmp180        Bmp180
	bmp280        Bmp280
	bmp388        Bmp388
	ds1407        Ds1407
	ds3231        Ds3231
	ft6336        Ft6336
	hd44780i2c    Hd44780i2c
	hts221        Hts221
	ina260        Ina260
	lis2mdl       Lis2mdl
	lis3dh        Lis3dh
	lps22hb       Lps22hb
	lsm6ds3       Lsm6ds3
	lsm6dsox      Lsm6dsox
	lsm9ds1       Lsm9ds1
	lsm303agr     Lsm303agr
	mag3110       Mag3110
	mma8653       Mma8653
	mpu6050       Mpu6050
	pcf8563       Pcf8563
	sht3x         Sht3x
	tmp102        Tmp102
	veml6070      Veml6070
	vl53l1xDevice Vl53l1xDevice
}

type device interface {
	configure()
	newDevice()
	connect()
}

func newDevice(d device) interface{} {
	d.configure()

	return nil
}
