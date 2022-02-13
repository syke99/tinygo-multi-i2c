package multi

import "time"

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
// BMP388
func (d Bmp388) readRegister(register byte, len int) (data []byte, err error) {
	data = make([]byte, len)
	err = d.bus.ReadRegister(d.Address, register, data)
	return
}

func (d Bmp388) writeRegister(register byte, data byte) error {
	return d.bus.WriteRegister(d.Address, register, []byte{data})
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
