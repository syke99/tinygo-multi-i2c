package multi

import (
	"errors"
	"time"
)

// TODO: flesh out configuration methods for each device

func (d Adxl345) configure() error {
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_BW_RATE, []byte{d.bwRate.toByte()})
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_POWER_CTL, []byte{d.powerCtl.toByte()})
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_DATA_FORMAT, []byte{d.dataFormat.toByte()})

	return nil
}

func (d Amg88xx) configure() error {
	d.data = make([]uint8, 128)

	d.SetPCTL(AMG88XX_NORMAL_MODE)
	d.SetReset(AMG88XX_INITIAL_RESET)
	d.SetFrameRate(AMG88XX_FPS_10)

	time.Sleep(100 * time.Millisecond)

	return nil
}

func (d At24cx) configure() {
	// At24cx may need to be removed?
}

func (d Bh1750) configure() error {
	if err := d.bus.Tx(d.Address, []byte{BH1750_POWER_ON}, nil); err != nil {
		return err
	}
	d.SetMode(d.mode)

	return nil
}

func (d Blinkm) configure() error {
	if err := d.bus.Tx(d.Address, []byte{'o'}, nil); err != nil {
		return err
	}

	return nil
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

func (d Bmp388) configure() {
	// Bmp388 may need to be removed?
}

func (d Ds3231) configure() {
	// DS3231 doesn't need a configure method??
}

func (d Ina260) configure() {
	// Ina260 may need to be removed?
}

func (d Lis3dh) configure() error {
	// enable all axes, normal mode
	if err := d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL1, []byte{0x07}); err != nil {
		return err
	}

	// 400Hz rate
	d.SetDataRate(LIS3DH_DATARATE_400_HZ)

	// High res & BDU enabled
	if err := d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL4, []byte{0x88}); err != nil {
		return err
	}

	// get current range
	d.r = d.ReadRange()

	return nil
}

func (d Lps22hb) configure() {
	// Lps22hb doesn't need a configure method??
}

func (d Mpu6050) configure() {
	d.bus.WriteRegister(uint8(d.Address), MPU6050_PWR_MGMT_1, []uint8{0})
}

func (d Sht3x) configure() {
	// Sht3x doesn't need a configure method??
}

func (d Vl53l1x) configure(use2v8Mode bool) error {
	if !d.Connected() {
		return errors.New("Vl53l1x device not connected.")
	}
	d.writeReg(VL53L1x_SOFT_RESET, 0x00)
	time.Sleep(100 * time.Microsecond)
	d.writeReg(VL53L1x_SOFT_RESET, 0x01)
	time.Sleep(1 * time.Millisecond)

	start := time.Now()
	for (d.readReg(VL53L1x_FIRMWARE_SYSTEM_STATUS) & 0x01) == 0 {
		elapsed := time.Since(start)
		if d.timeout > 0 && uint32(elapsed.Seconds()*1000) > d.timeout {
			return errors.New("Connection timed out.")
		}
	}

	if use2v8Mode {
		d.writeReg(VL53L1x_PAD_I2C_HV_EXTSUP_CONFIG, d.readReg(VL53L1x_PAD_I2C_HV_EXTSUP_CONFIG)|0x01)
	}

	d.fastOscillatorFreq = d.readReg16Bit(VL53L1x_OSC_MEASURED_FAST_OSC_FREQUENCY)
	d.oscillatorOffset = d.readReg16Bit(VL53L1x_RESULT_OSC_CALIBRATE_VAL)

	// static config
	d.writeReg16Bit(VL53L1x_DSS_CONFIG_TARGET_TOTAL_RATE_MCPS, VL53L1x_TARGETRATE)
	d.writeReg(VL53L1x_GPIO_TIO_HV_STATUS, 0x02)
	d.writeReg(VL53L1x_SIGMA_ESTIMATOR_EFFECTIVE_PULSE_WIDTH_NS, 8)
	d.writeReg(VL53L1x_SIGMA_ESTIMATOR_EFFECTIVE_AMBIENT_WIDTH_NS, 16)
	d.writeReg(VL53L1x_ALGO_CROSSTALK_COMPENSATION_VALID_HEIGHT_MM, 0xFF)
	d.writeReg(VL53L1x_ALGO_RANGE_MIN_CLIP, 0)
	d.writeReg(VL53L1x_ALGO_CONSISTENCY_CHECK_TOLERANCE, 2)

	// general config
	d.writeReg16Bit(VL53L1x_SYSTEM_THRESH_RATE_HIGH, 0x0000)
	d.writeReg16Bit(VL53L1x_SYSTEM_THRESH_RATE_LOW, 0x0000)
	d.writeReg(VL53L1x_DSS_CONFIG_APERTURE_ATTENUATION, 0x38)

	// timing config
	d.writeReg16Bit(VL53L1x_RANGE_CONFIG_SIGMA_THRESH, 360)
	d.writeReg16Bit(VL53L1x_RANGE_CONFIG_MIN_COUNT_RATE_RTN_LIMIT_MCPS, 192)

	// dynamic config
	d.writeReg(VL53L1x_SYSTEM_GROUPED_PARAMETER_HOLD_0, 0x01)
	d.writeReg(VL53L1x_SYSTEM_GROUPED_PARAMETER_HOLD_1, 0x01)
	d.writeReg(VL53L1x_SD_CONFIG_QUANTIFIER, 2)

	d.writeReg(VL53L1x_SYSTEM_GROUPED_PARAMETER_HOLD, 0x00)
	d.writeReg(VL53L1x_SYSTEM_SEED_CONFIG, 1)

	// Low power auto mode
	d.writeReg(VL53L1x_SYSTEM_SEQUENCE_CONFIG, 0x8B) // VHV, PHASECAL, DSS1, RANGE
	d.writeReg16Bit(VL53L1x_DSS_CONFIG_MANUAL_EFFECTIVE_SPADS_SELECT, 200<<8)
	d.writeReg(VL53L1x_DSS_CONFIG_ROI_MODE_CONTROL, 2) // REQUESTED_EFFFECTIVE_SPADS

	d.SetDistanceMode(d.mode)
	d.SetMeasurementTimingBudget(50000)

	d.writeReg16Bit(VL53L1x_ALGO_PART_TO_PART_RANGE_OFFSET_MM, d.readReg16Bit(VL53L1x_MM_CONFIG_OUTER_OFFSET_MM)*4)

	return nil
}

func timeoutMicrosecondsToMclks(timeoutMicroseconds uint32, macroPeriodMicroseconds uint32) uint32 {
	return ((timeoutMicroseconds << 12) + (macroPeriodMicroseconds >> 1)) / macroPeriodMicroseconds
}

func timeoutMclksToMicroseconds(timeoutMclks uint32, macroPeriodMicroseconds uint32) uint32 {
	return uint32((uint64(timeoutMclks)*uint64(macroPeriodMicroseconds) + 0x800) >> 12)
}

func encodeTimeout(timeoutMclks uint32) uint16 {
	if timeoutMclks == 0 {
		return 0
	}
	msb := 0
	lsb := timeoutMclks - 1
	for (lsb & 0xFFFFFF00) > 0 {
		lsb >>= 1
		msb++
	}
	return uint16(msb<<8) | uint16(lsb&0xFF)
}

// decodeTimeout decodes the timeout from the format: (LSByte * 2^MSByte) + 1
func decodeTimeout(regVal uint16) uint32 {
	return (uint32(regVal&0xFF) << (regVal >> 8)) + 1
}

func readUint(msb byte, lsb byte) uint16 {
	return (uint16(msb) << 8) | uint16(lsb)
}

func readUint32(data []byte) uint32 {
	if len(data) != 4 {
		return 0
	}
	var value uint32
	value = uint32(data[0]) << 24
	value |= uint32(data[1]) << 16
	value |= uint32(data[2]) << 8
	value |= uint32(data[3])
	return value
}

func readUintLE(msb byte, lsb byte) uint16 {
	temp := readUint(msb, lsb)
	return (temp >> 8) | (temp << 8)
}

// readIntLE converts two little endian bytes to int16
func readIntLE(msb byte, lsb byte) int16 {
	return int16(readUintLE(msb, lsb))
}
