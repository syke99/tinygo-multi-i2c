package multi

import "time"

// TODO: flesh out configuration methods for each device

func (d Adxl345) configure() {
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_BW_RATE, []byte{d.bwRate.toByte()})
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_POWER_CTL, []byte{d.powerCtl.toByte()})
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_DATA_FORMAT, []byte{d.dataFormat.toByte()})
}

func (d Amg88xx) configure() {
	d.data = make([]uint8, 128)

	d.SetPCTL(AMG88XX_NORMAL_MODE)
	d.SetReset(AMG88XX_INITIAL_RESET)
	d.SetFrameRate(AMG88XX_FPS_10)

	time.Sleep(100 * time.Millisecond)
}

func (d At24cx) configure() {

}

func (d Bh1750) configure() {
	d.bus.Tx(d.Address, []byte{BH1750_POWER_ON}, nil)
	d.SetMode(d.mode)
}

func (d Blinkm) configure() {
	d.bus.Tx(d.Address, []byte{'o'}, nil)
}

func (d Bme280) configure() {
	var data [24]byte
	err := d.bus.ReadRegister(uint8(d.Address), BME280_REG_CALIBRATION, data[:])
	if err != nil {
		return
	}

	var h1 [1]byte
	err = d.bus.ReadRegister(uint8(d.Address), BME280_REG_CALIBRATION_H1, h1[:])
	if err != nil {
		return
	}

	var h2lsb [7]byte
	err = d.bus.ReadRegister(uint8(d.Address), BME280_REG_CALIBRATION_H2LSB, h2lsb[:])
	if err != nil {
		return
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
}

func (d Bmp280) Configure(standby bmp290Standby, filter bmp280Filter, temp bmp280Oversampling, pres bmp280Oversampling, mode bmp280Mode) {
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
		return
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
}

func (d Bmp388) configure() {

}

func (d Ds3231) configure() {
	// DS3231 doesn't need a configure method??
}

func (d Ina260) configure() {

}

func (d Lis3dh) configure() {
	// enable all axes, normal mode
	d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL1, []byte{0x07})

	// 400Hz rate
	d.SetDataRate(LIS3DH_DATARATE_400_HZ)

	// High res & BDU enabled
	d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL4, []byte{0x88})

	// get current range
	d.r = d.ReadRange()
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

func (d Vl53l1x) Configure(use2v8Mode bool) bool {
	if !d.Connected() {
		return false
	}
	d.writeReg(VL53L1x_SOFT_RESET, 0x00)
	time.Sleep(100 * time.Microsecond)
	d.writeReg(VL53L1x_SOFT_RESET, 0x01)
	time.Sleep(1 * time.Millisecond)

	start := time.Now()
	for (d.readReg(VL53L1x_FIRMWARE_SYSTEM_STATUS) & 0x01) == 0 {
		elapsed := time.Since(start)
		if d.timeout > 0 && uint32(elapsed.Seconds()*1000) > d.timeout {
			return false
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

	return true
}

//-------------------------------------------------------------------------------------
// ADXL345
func (p adxl345PowerCtl) toByte() (bits uint8) {
	bits = 0x00
	bits = bits | (p.link << 5)
	bits = bits | (p.autoSleep << 4)
	bits = bits | (p.measure << 3)
	bits = bits | (p.sleep << 2)
	bits = bits | p.wakeUp

	return bits
}

// toByte returns a byte from the dataFormat configuration
func (d adxl345DataFormat) toByte() (bits uint8) {
	bits = 0x00
	bits = bits | (d.selfTest << 7)
	bits = bits | (d.spi << 6)
	bits = bits | (d.intInvert << 5)
	bits = bits | (d.fullRes << 3)
	bits = bits | (d.justify << 2)
	bits = bits | uint8(d.sensorRange)

	return bits
}

// toByte returns a byte from the bwRate configuration
func (b adxl345BwRate) toByte() (bits uint8) {
	bits = 0x00
	bits = bits | (b.lowPower << 4)
	bits = bits | uint8(b.rate)

	return bits
}

//-------------------------------------------------------------------------------------
// AMG88XX
func (d Amg88xx) SetPCTL(pctl uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_PCTL, []byte{pctl})
}

// SetReset sets the reset value
func (d Amg88xx) SetReset(rst uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_RST, []byte{rst})
}

// SetFrameRate configures the frame rate
func (d Amg88xx) SetFrameRate(framerate uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_FPSC, []byte{framerate & 0x01})
}

//-------------------------------------------------------------------------------------
// BH170
// SetMode changes the reading mode for the sensor
func (d Bh1750) SetMode(mode byte) {
	d.mode = mode
	d.bus.Tx(d.Address, []byte{byte(d.mode)}, nil)
	time.Sleep(10 * time.Millisecond)
}

//-------------------------------------------------------------------------------------
// Lis3dh
func (d Lis3dh) SetDataRate(rate Lis3dhDataRate) {
	ctl1 := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), LIS3DH_REG_CTRL1, ctl1)
	if err != nil {
		println(err.Error())
	}
	// mask off bits
	ctl1[0] &^= 0xf0
	ctl1[0] |= (byte(rate) << 4)
	d.bus.WriteRegister(uint8(d.Address), LIS3DH_REG_CTRL1, ctl1)
}

func (d Lis3dh) ReadRange() (r Lis3dhRange) {
	ctl := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), LIS3DH_REG_CTRL4, ctl)
	if err != nil {
		println(err.Error())
	}
	// mask off bits
	r = Lis3dhRange(ctl[0] >> 4)
	r &= 0x03

	return r
}

//-------------------------------------------------------------------------------------
// VL53L1x
func (d Vl53l1x) Connected() bool {
	return d.readReg16Bit(VL53L1x_WHO_AM_I) == VL53L1x_CHIP_ID
}

// writeReg sends a single byte to the specified register address
func (d Vl53l1x) writeReg(reg uint16, value uint8) {
	msb := byte((reg >> 8) & 0xFF)
	lsb := byte(reg & 0xFF)
	d.bus.Tx(d.Address, []byte{msb, lsb, value}, nil)
}

// writeReg16Bit sends two bytes to the specified register address
func (d Vl53l1x) writeReg16Bit(reg uint16, value uint16) {
	data := make([]byte, 4)
	data[0] = byte((reg >> 8) & 0xFF)
	data[1] = byte(reg & 0xFF)
	data[2] = byte((value >> 8) & 0xFF)
	data[3] = byte(value & 0xFF)
	d.bus.Tx(d.Address, data, nil)
}

// writeReg32Bit sends four bytes to the specified register address
func (d Vl53l1x) writeReg32Bit(reg uint16, value uint32) {
	data := make([]byte, 6)
	data[0] = byte((reg >> 8) & 0xFF)
	data[1] = byte(reg & 0xFF)
	data[2] = byte((value >> 24) & 0xFF)
	data[3] = byte((value >> 16) & 0xFF)
	data[4] = byte((value >> 8) & 0xFF)
	data[5] = byte(value & 0xFF)
	d.bus.Tx(d.Address, data, nil)
}

// readReg reads a single byte from the specified address
func (d Vl53l1x) readReg(reg uint16) uint8 {
	data := []byte{0}
	msb := byte((reg >> 8) & 0xFF)
	lsb := byte(reg & 0xFF)
	d.bus.Tx(d.Address, []byte{msb, lsb}, data)
	return data[0]
}

// readReg16Bit reads two bytes from the specified address
// and returns it as a uint16
func (d Vl53l1x) readReg16Bit(reg uint16) uint16 {
	data := []byte{0, 0}
	msb := byte((reg >> 8) & 0xFF)
	lsb := byte(reg & 0xFF)
	d.bus.Tx(d.Address, []byte{msb, lsb}, data)
	return readUint(data[0], data[1])
}

// readReg32Bit reads four bytes from the specified address
// and returns it as a uint32
func (d Vl53l1x) readReg32Bit(reg uint16) uint32 {
	data := make([]byte, 4)
	msb := byte((reg >> 8) & 0xFF)
	lsb := byte(reg & 0xFF)
	d.bus.Tx(d.Address, []byte{msb, lsb}, data)
	return readUint32(data)
}

func (d Vl53l1x) SetMeasurementTimingBudget(budgetMicroseconds uint32) bool {
	if budgetMicroseconds <= VL53L1x_TIMING_GUARD {
		return false
	}
	budgetMicroseconds -= VL53L1x_TIMING_GUARD
	if budgetMicroseconds > 1100000 {
		return false
	}
	rangeConfigTimeout := budgetMicroseconds / 2
	// Update Macro Period for Range A VCSEL Period
	macroPeriod := d.calculateMacroPeriod(uint32(d.readReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A)))

	// Update Phase timeout - uses Timing A
	phasecalTimeoutMclks := timeoutMicrosecondsToMclks(1000, macroPeriod)
	if phasecalTimeoutMclks > 0xFF {
		phasecalTimeoutMclks = 0xFF
	}
	d.writeReg(VL53L1x_PHASECAL_CONFIG_TIMEOUT_MACROP, uint8(phasecalTimeoutMclks))

	// Update MM Timing A timeout
	d.writeReg16Bit(VL53L1x_MM_CONFIG_TIMEOUT_MACROP_A, encodeTimeout(timeoutMicrosecondsToMclks(1, macroPeriod)))
	// Update Range Timing A timeout
	d.writeReg16Bit(VL53L1x_RANGE_CONFIG_TIMEOUT_MACROP_A, encodeTimeout(timeoutMicrosecondsToMclks(rangeConfigTimeout, macroPeriod)))

	macroPeriod = d.calculateMacroPeriod(uint32(d.readReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A)))
	// Update MM Timing B timeout
	d.writeReg16Bit(VL53L1x_MM_CONFIG_TIMEOUT_MACROP_B, encodeTimeout(timeoutMicrosecondsToMclks(1, macroPeriod)))
	// Update Range Timing B timeout
	d.writeReg16Bit(VL53L1x_RANGE_CONFIG_TIMEOUT_MACROP_B, encodeTimeout(timeoutMicrosecondsToMclks(rangeConfigTimeout, macroPeriod)))

	return true
}

func (d Vl53l1x) calculateMacroPeriod(vcselPeriod uint32) uint32 {
	pplPeriodMicroseconds := (uint32(1) << 30) / uint32(d.fastOscillatorFreq)
	vcselPeriodPclks := (vcselPeriod + 1) << 1
	macroPeriodMicroseconds := 2304 * pplPeriodMicroseconds
	macroPeriodMicroseconds >>= 6
	macroPeriodMicroseconds *= vcselPeriodPclks
	macroPeriodMicroseconds >>= 6
	return macroPeriodMicroseconds
}

func (d Vl53l1x) SetDistanceMode(mode vl53l1xDistanceMode) bool {
	budgetMicroseconds := d.GetMeasurementTimingBudget()
	switch mode {
	case VL53L1x_SHORT:
		// timing config
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A, 0x07)
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_B, 0x05)
		d.writeReg(VL53L1x_RANGE_CONFIG_VALID_PHASE_HIGH, 0x38)

		// dynamic config
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD0, 0x07)
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD1, 0x05)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD0, 6)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD1, 6)
		break
	case VL53L1x_MEDIUM:
		// timing config
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A, 0x0B)
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_B, 0x09)
		d.writeReg(VL53L1x_RANGE_CONFIG_VALID_PHASE_HIGH, 0x78)

		// dynamic config
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD0, 0x0B)
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD1, 0x09)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD0, 10)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD1, 10)
		break
	case VL53L1x_LONG:
		// timing config
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A, 0x0F)
		d.writeReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_B, 0x0D)
		d.writeReg(VL53L1x_RANGE_CONFIG_VALID_PHASE_HIGH, 0xB8)

		// dynamic config
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD0, 0x0F)
		d.writeReg(VL53L1x_SD_CONFIG_WOI_SD1, 0x0D)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD0, 14)
		d.writeReg(VL53L1x_SD_CONFIG_INITIAL_PHASE_SD1, 14)
		break
	default:
		return false
	}

	d.SetMeasurementTimingBudget(budgetMicroseconds)
	d.mode = mode
	return true
}

func (d Vl53l1x) GetMeasurementTimingBudget() uint32 {
	macroPeriod := d.calculateMacroPeriod(uint32(d.readReg(VL53L1x_RANGE_CONFIG_VCSEL_PERIOD_A)))
	rangeConfigTimeout := timeoutMclksToMicroseconds(decodeTimeout(d.readReg16Bit(VL53L1x_RANGE_CONFIG_TIMEOUT_MACROP_A)), macroPeriod)
	return 2 * uint32(rangeConfigTimeout) * VL53L1x_TIMING_GUARD
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
