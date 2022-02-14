package multi

type Devices struct {
	adxl345 Adxl345
	amg88xx Amg88xx
	at24cx  At24cx
	bh1750  Bh1750
	blinkm  Blinkm
	bme280  Bme280
	bmp280  Bmp280
	bmp388  Bmp388
	ds3231  Ds3231
	ina260  Ina260
	lis3dh  Lis3dh
	lps22hb Lps22hb
	lsm6ds3 Lsm6ds3
	mpu6050 Mpu6050
	sht3x   Sht3x
	vl53l1x Vl53l1x
}

func (d *Devices) addAdxl1345(v Adxl345) {
	d.adxl345 = v
}

func (d *Devices) addAmg88xx(v Amg88xx) {
	d.amg88xx = v
}

func (d *Devices) addAt24cx(v At24cx) {
	d.at24cx = v
}

func (d *Devices) addBh1750(v Bh1750) {
	d.bh1750 = v
}

func (d *Devices) addBlinkm(v Blinkm) {
	d.blinkm = v
}

func (d *Devices) addBme280(v Bme280) {
	d.bme280 = v
}

func (d *Devices) addBmp280(v Bmp280) {
	d.bmp280 = v
}

func (d *Devices) addBmp388(v Bmp388) {
	d.bmp388 = v
}

func (d *Devices) addDs3231(v Ds3231) {
	d.ds3231 = v
}

func (d *Devices) addIna260(v Ina260) {
	d.ina260 = v
}

func (d *Devices) addLis3dh(v Lis3dh) {
	d.lis3dh = v
}

func (d *Devices) addLps22hb(v Lps22hb) {
	d.lps22hb = v
}

func (d *Devices) addLsm6ds3(v Lsm6ds3) {
	d.lsm6ds3 = v
}

func (d *Devices) addMpu6050(v Mpu6050) {
	d.mpu6050 = v
}

func (d *Devices) addSht3x(v Sht3x) {
	d.sht3x = v
}

func (d *Devices) addVl53l1x(v Vl53l1x) {
	d.vl53l1x = v
}
