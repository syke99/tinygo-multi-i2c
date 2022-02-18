package v2

type Devices struct {
	adxl345 []Adxl345
	amg88xx []Amg88xx
	bh1750  []Bh1750
	blinkm  []Blinkm
	bme280  []Bme280
	bmp280  []Bmp280
	lis3dh  []Lis3dh
	lps22hb []Lps22hb
	mpu6050 []Mpu6050
}

func (d *Devices) addAdxl1345(v Adxl345) {
	d.adxl345 = append(d.adxl345, v)
}

func (d *Devices) addAmg88xx(v Amg88xx) {
	d.amg88xx = append(d.amg88xx, v)
}

func (d *Devices) addBh1750(v Bh1750) {
	d.bh1750 = append(d.bh1750, v)
}

func (d *Devices) addBlinkm(v Blinkm) {
	d.blinkm = append(d.blinkm, v)
}

func (d *Devices) addBme280(v Bme280) {
	d.bme280 = append(d.bme280, v)
}

func (d *Devices) addBmp280(v Bmp280) {
	d.bmp280 = append(d.bmp280, v)
}

func (d *Devices) addLis3dh(v Lis3dh) {
	d.lis3dh = append(d.lis3dh, v)
}

func (d *Devices) addLps22hb(v Lps22hb) {
	d.lps22hb = append(d.lps22hb, v)
}

func (d *Devices) addMpu6050(v Mpu6050) {
	d.mpu6050 = append(d.mpu6050, v)
}
