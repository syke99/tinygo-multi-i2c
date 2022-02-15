package multi

type Devices struct {
	Adxl345
	Amg88xx
	Bh1750
	Blinkm
	Bme280
	Bmp280
	Ds3231
	Lis3dh
	Lps22hb
	Mpu6050
}

func (d *Devices) addAdxl1345(v Adxl345) {
	d.Adxl345 = v
}

func (d *Devices) addAmg88xx(v Amg88xx) {
	d.Amg88xx = v
}

func (d *Devices) addBh1750(v Bh1750) {
	d.Bh1750 = v
}

func (d *Devices) addBlinkm(v Blinkm) {
	d.Blinkm = v
}

func (d *Devices) addBme280(v Bme280) {
	d.Bme280 = v
}

func (d *Devices) addBmp280(v Bmp280) {
	d.Bmp280 = v
}

func (d *Devices) addDs3231(v Ds3231) {
	d.Ds3231 = v
}

func (d *Devices) addLis3dh(v Lis3dh) {
	d.Lis3dh = v
}

func (d *Devices) addLps22hb(v Lps22hb) {
	d.Lps22hb = v
}

func (d *Devices) addMpu6050(v Mpu6050) {
	d.Mpu6050 = v
}
