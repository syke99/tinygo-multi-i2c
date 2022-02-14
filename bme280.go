package multi

import (
	"math"

	"tinygo.org/x/drivers"
)

type bme280CalibrationCoefficients struct {
	t1 uint16
	t2 int16
	t3 int16
	p1 uint16
	p2 int16
	p3 int16
	p4 int16
	p5 int16
	p6 int16
	p7 int16
	p8 int16
	p9 int16
	h1 uint8
	h2 int16
	h3 uint8
	h4 int16
	h5 int16
	h6 int8
}

type Bme280 struct {
	bus                     drivers.I2C
	Address                 uint16
	calibrationCoefficients bme280CalibrationCoefficients
}

func newBme280(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Bme280{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Bme280{
			bus:     bus,
			Address: Bme280Address,
		}
	}
}

func (d Bme280) configure() error {
	var data [24]byte
	err := d.bus.ReadRegister(uint8(d.Address), BME280_REG_CALIBRATION, data[:])
	if err != nil {
		return err
	}

	var h1 [1]byte
	err = d.bus.ReadRegister(uint8(d.Address), BME280_REG_CALIBRATION_H1, h1[:])
	if err != nil {
		return err
	}

	var h2lsb [7]byte
	err = d.bus.ReadRegister(uint8(d.Address), BME280_REG_CALIBRATION_H2LSB, h2lsb[:])
	if err != nil {
		return err
	}

	d.calibrationCoefficients.t1 = readUintLE(data[0], data[1])
	d.calibrationCoefficients.t2 = readIntLE(data[2], data[3])
	d.calibrationCoefficients.t3 = readIntLE(data[4], data[5])
	d.calibrationCoefficients.p1 = readUintLE(data[6], data[7])
	d.calibrationCoefficients.p2 = readIntLE(data[8], data[9])
	d.calibrationCoefficients.p3 = readIntLE(data[10], data[11])
	d.calibrationCoefficients.p4 = readIntLE(data[12], data[13])
	d.calibrationCoefficients.p5 = readIntLE(data[14], data[15])
	d.calibrationCoefficients.p6 = readIntLE(data[16], data[17])
	d.calibrationCoefficients.p7 = readIntLE(data[18], data[19])
	d.calibrationCoefficients.p8 = readIntLE(data[20], data[21])
	d.calibrationCoefficients.p9 = readIntLE(data[22], data[23])

	d.calibrationCoefficients.h1 = h1[0]
	d.calibrationCoefficients.h2 = readIntLE(h2lsb[0], h2lsb[1])
	d.calibrationCoefficients.h3 = h2lsb[2]
	d.calibrationCoefficients.h6 = int8(h2lsb[6])
	d.calibrationCoefficients.h4 = 0 + (int16(h2lsb[3]) << 4) | (int16(h2lsb[4] & 0x0F))
	d.calibrationCoefficients.h5 = 0 + (int16(h2lsb[5]) << 4) | (int16(h2lsb[4]) >> 4)

	d.bus.WriteRegister(uint8(d.Address), BME280_CTRL_HUMIDITY_ADDR, []byte{0x3f})
	d.bus.WriteRegister(uint8(d.Address), BME280_CTRL_MEAS_ADDR, []byte{0xB7})
	d.bus.WriteRegister(uint8(d.Address), BME280_CTRL_CONFIG, []byte{0x00})

	return nil
}

func (d Bme280) connected() bool {
	data := []byte{0}
	d.bus.ReadRegister(uint8(d.Address), BME280_WHO_AM_I, data)
	return data[0] == BME280_CHIP_ID
}

func (d Bme280) Reset() {
	d.bus.WriteRegister(uint8(d.Address), BME280_CMD_RESET, []byte{0xB6})
}

func (d Bme280) ReadTemperature() (int32, error) {
	data, err := d.readData()
	if err != nil {
		return 0, err
	}

	temp, _ := d.calculateTemp(data)
	return temp, nil
}

func (d Bme280) ReadPressure() (int32, error) {
	data, err := d.readData()
	if err != nil {
		return 0, err
	}
	_, tFine := d.calculateTemp(data)
	pressure := d.calculatePressure(data, tFine)
	return pressure, nil
}

func (d Bme280) ReadHumidity() (int32, error) {
	data, err := d.readData()
	if err != nil {
		return 0, err
	}
	_, tFine := d.calculateTemp(data)
	humidity := d.calculateHumidity(data, tFine)
	return humidity, nil
}

func (d Bme280) ReadAltitude() (alt int32, err error) {
	mPa, _ := d.ReadPressure()
	atmP := float32(mPa) / 100000
	alt = int32(44330.0 * (1.0 - math.Pow(float64(atmP/BME280_SEALEVEL_PRESSURE), 0.1903)))
	return
}

// private functions

func (d Bme280) calculateTemp(data [8]byte) (int32, int32) {

	rawTemp := convert3Bytes(data[3], data[4], data[5])

	var1 := (((rawTemp >> 3) - (int32(d.calibrationCoefficients.t1) << 1)) * int32(d.calibrationCoefficients.t2)) >> 11
	var2 := (((((rawTemp >> 4) - int32(d.calibrationCoefficients.t1)) * ((rawTemp >> 4) - int32(d.calibrationCoefficients.t1))) >> 12) * int32(d.calibrationCoefficients.t3)) >> 14

	tFine := var1 + var2
	T := (tFine*5 + 128) >> 8
	return (10 * T), tFine
}

func (d Bme280) readData() (data [8]byte, err error) {
	err = d.bus.ReadRegister(uint8(d.Address), BME280_REG_PRESSURE, data[:])
	if err != nil {
		println(err)
		return
	}
	return
}

func (d Bme280) calculatePressure(data [8]byte, tFine int32) int32 {

	rawPressure := convert3Bytes(data[0], data[1], data[2])

	var1 := int64(tFine) - 128000
	var2 := var1 * var1 * int64(d.calibrationCoefficients.p6)
	var2 = var2 + ((var1 * int64(d.calibrationCoefficients.p5)) << 17)
	var2 = var2 + (int64(d.calibrationCoefficients.p4) << 35)
	var1 = ((var1 * var1 * int64(d.calibrationCoefficients.p3)) >> 8) + ((var1 * int64(d.calibrationCoefficients.p2)) << 12)
	var1 = ((int64(1) << 47) + var1) * int64(d.calibrationCoefficients.p1) >> 33

	if var1 == 0 {
		return 0 // avoid exception caused by division by zero
	}
	p := int64(1048576 - rawPressure)
	p = (((p << 31) - var2) * 3125) / var1
	var1 = (int64(d.calibrationCoefficients.p9) * (p >> 13) * (p >> 13)) >> 25
	var2 = (int64(d.calibrationCoefficients.p8) * p) >> 19

	p = ((p + var1 + var2) >> 8) + (int64(d.calibrationCoefficients.p7) << 4)
	p = (p / 256)
	return int32(1000 * p)
}

func (d Bme280) calculateHumidity(data [8]byte, tFine int32) int32 {

	rawHumidity := convert2Bytes(data[6], data[7])

	h := float32(tFine) - 76800

	if h == 0 {
		println("invalid value")
	}

	var1 := float32(rawHumidity) - (float32(d.calibrationCoefficients.h4)*64.0 +
		(float32(d.calibrationCoefficients.h5) / 16384.0 * h))

	var2 := float32(d.calibrationCoefficients.h2) / 65536.0 *
		(1.0 + float32(d.calibrationCoefficients.h6)/67108864.0*h*
			(1.0+float32(d.calibrationCoefficients.h3)/67108864.0*h))

	h = var1 * var2
	h = h * (1 - float32(d.calibrationCoefficients.h1)*h/524288)
	return int32(100 * h)

}

func convert2Bytes(msb byte, lsb byte) int32 {
	return int32(readUint(msb, lsb))
}

func convert3Bytes(msb byte, b1 byte, lsb byte) int32 {
	return int32(((((uint32(msb) << 8) | uint32(b1)) << 8) | uint32(lsb)) >> 4)
}
