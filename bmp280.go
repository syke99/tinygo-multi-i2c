package multi

import (
	"time"

	"tinygo.org/x/drivers"
)

//-------------------------------------------------------------------------------------
// Bmp280
type bmp280CalibrationCoefficients struct {
	// Temperature compensation
	t1 uint16
	t2 int16
	t3 int16

	// Pressure compensation
	p1 uint16
	p2 int16
	p3 int16
	p4 int16
	p5 int16
	p6 int16
	p7 int16
	p8 int16
	p9 int16
}

type bmp280Oversampling uint
type bmp280Mode uint
type bmp290Standby uint
type bmp280Filter uint

type Bmp280 struct {
	bus         drivers.I2C
	Address     uint16
	cali        bmp280CalibrationCoefficients
	Temperature bmp280Oversampling
	Pressure    bmp280Oversampling
	Mode        bmp280Mode
	Standby     bmp290Standby
	Filter      bmp280Filter
}

//-------------------------------------------------------------------------------------
// BMP280
func newBmp280(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Bmp280{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Bmp280{
			bus:     bus,
			Address: Bmp280Address,
		}
	}
}

//-------------------------------------------------------------------------------------
// Bmp280
func (d Bmp280) configure(standby bmp290Standby, filter bmp280Filter, temp bmp280Oversampling, pres bmp280Oversampling, mode bmp280Mode) error {
	d.Standby = standby
	d.Filter = filter
	d.Temperature = temp
	d.Pressure = pres
	d.Mode = mode

	//  Write the configuration (standby, filter, spi 3 wire)
	config := uint(d.Standby<<5) | uint(d.Filter<<2) | 0x00
	d.bus.WriteRegister(uint8(d.Address), BMP280_REG_CONFIG, []byte{byte(config)})

	// Write the control (temperature oversampling, pressure oversampling,
	config = uint(d.Temperature<<5) | uint(d.Pressure<<2) | uint(d.Mode)
	d.bus.WriteRegister(uint8(d.Address), BMP280_REG_CTRL_MEAS, []byte{byte(config)})

	// Read Calibration data
	data := make([]byte, 24)
	err := d.bus.ReadRegister(uint8(d.Address), BMP280_REG_CALI, data)
	if err != nil {
		return err
	}

	// Datasheet: 3.11.2 Trimming parameter readout
	d.cali.t1 = readUintLE(data[0], data[1])
	d.cali.t2 = readIntLE(data[2], data[3])
	d.cali.t3 = readIntLE(data[4], data[5])

	d.cali.p1 = readUintLE(data[6], data[7])
	d.cali.p2 = readIntLE(data[8], data[9])
	d.cali.p3 = readIntLE(data[10], data[11])
	d.cali.p4 = readIntLE(data[12], data[13])
	d.cali.p5 = readIntLE(data[14], data[15])
	d.cali.p6 = readIntLE(data[16], data[17])
	d.cali.p7 = readIntLE(data[18], data[19])
	d.cali.p8 = readIntLE(data[20], data[21])
	d.cali.p9 = readIntLE(data[22], data[23])

	return nil
}

//-------------------------------------------------------------------------------------
// BMP280
func (d Bmp280) connected() bool {
	data := make([]byte, 1)
	d.bus.ReadRegister(uint8(d.Address), BMP280_REG_ID, data)
	return data[0] == BMP280_CHIP_ID
}

//-------------------------------------------------------------------------------------
// BMP280
func (d Bmp280) Reset() {
	d.bus.WriteRegister(uint8(d.Address), BMP280_REG_RESET, []byte{BMP280_CMD_RESET})
}

func (d Bmp280) ReadTemperature() (temperature int32, err error) {
	data, err := d.readData(BMP280_REG_TEMP, 3)
	if err != nil {
		return
	}

	rawTemp := convert3Bytes(data[0], data[1], data[2])

	// Datasheet: 8.2 Compensation formula in 32 bit fixed point
	// Temperature compensation
	var1 := ((rawTemp >> 3) - int32(d.cali.t1<<1)) * int32(d.cali.t2) >> 11
	var2 := (((rawTemp >> 4) - int32(d.cali.t1)) * ((rawTemp >> 4) - int32(d.cali.t1)) >> 12) *
		int32(d.cali.t3) >> 14

	tFine := var1 + var2

	// Convert from degrees to milli degrees by multiplying by 10.
	// Will output 30250 milli degrees celsius for 30.25 degrees celsius
	temperature = 10 * ((tFine*5 + 128) >> 8)
	return
}

func (d Bmp280) ReadPressure() (pressure int32, err error) {
	// First 3 bytes are Pressure, last 3 bytes are Temperature
	data, err := d.readData(BMP280_REG_PRES, 6)
	if err != nil {
		return
	}

	rawTemp := convert3Bytes(data[3], data[4], data[5])

	// Datasheet: 8.2 Compensation formula in 32 bit fixed point
	// Calculate tFine (temperature), used for the Pressure compensation
	var1 := ((rawTemp >> 3) - int32(d.cali.t1<<1)) * int32(d.cali.t2) >> 11
	var2 := (((rawTemp >> 4) - int32(d.cali.t1)) * ((rawTemp >> 4) - int32(d.cali.t1)) >> 12) *
		int32(d.cali.t3) >> 14

	tFine := var1 + var2

	rawPres := convert3Bytes(data[0], data[1], data[2])

	// Datasheet: 8.2 Compensation formula in 32 bit fixed point
	// Pressure compensation
	var1 = (tFine >> 1) - 64000
	var2 = (((var1 >> 2) * (var1 >> 2)) >> 11) * int32(d.cali.p6)
	var2 = var2 + ((var1 * int32(d.cali.p5)) << 1)
	var2 = (var2 >> 2) + (int32(d.cali.p4) << 16)
	var1 = (((int32(d.cali.p3) * (((var1 >> 2) * (var1 >> 2)) >> 13)) >> 3) +
		((int32(d.cali.p2) * var1) >> 1)) >> 18
	var1 = ((32768 + var1) * int32(d.cali.p1)) >> 15

	if var1 == 0 {
		return 0, nil
	}

	p := uint32(((1048576 - rawPres) - (var2 >> 12)) * 3125)
	if p < 0x80000000 {
		p = (p << 1) / uint32(var1)
	} else {
		p = (p / uint32(var1)) * 2
	}

	var1 = (int32(d.cali.p9) * int32(((p>>3)*(p>>3))>>13)) >> 12
	var2 = (int32(p>>2) * int32(d.cali.p8)) >> 13

	return 1000 * (int32(p) + ((var1 + var2 + int32(d.cali.p7)) >> 4)), nil
}

func (d Bmp280) readData(register int, n int) ([]byte, error) {
	// If not in normal mode, set the mode to FORCED mode, to prevent incorrect measurements
	// After the measurement in FORCED mode, the sensor will return to SLEEP mode
	if d.Mode != BMP280_MODE_NORMAL {
		config := uint(d.Temperature<<5) | uint(d.Pressure<<2) | uint(BMP280_MODE_FORCED)
		d.bus.WriteRegister(uint8(d.Address), BMP280_REG_CTRL_MEAS, []byte{byte(config)})
	}

	// Check STATUS register, wait if data is not available yet
	status := make([]byte, 1)
	for d.bus.ReadRegister(uint8(d.Address), uint8(BMP280_REG_STATUS), status[0:]); status[0] != 4 && status[0] != 0; d.bus.ReadRegister(uint8(d.Address), uint8(BMP280_REG_STATUS), status[0:]) {
		time.Sleep(time.Millisecond)
	}

	// Read the requested register
	data := make([]byte, n)
	err := d.bus.ReadRegister(uint8(d.Address), uint8(register), data[:])
	return data, err
}
