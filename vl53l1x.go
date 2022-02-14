package multi

import (
	"errors"
	"time"

	"tinygo.org/x/drivers"
)

//-------------------------------------------------------------------------------------
// VL53L1x
type vl53l1xDistanceMode uint8
type vl53l1xRangeStatus uint8

type vl53l1xRangingData struct {
	mm              uint16
	status          vl53l1xRangeStatus
	signalRateMCPS  int32
	ambientRateMCPS int32
}

type vl53l1xResultBuffer struct {
	status                     uint8
	streamCount                uint8
	effectiveSPADCount         uint16
	ambientRateMCPSSD0         uint16
	mmCrosstalkSD0             uint16
	signalRateCrosstalkMCPSSD0 uint16
}

type Vl53l1x struct {
	bus                drivers.I2C
	Address            uint16
	mode               vl53l1xDistanceMode
	timeout            uint32
	fastOscillatorFreq uint16
	oscillatorOffset   uint16
	calibrated         bool
	VHVInit            uint8
	VHVTimeout         uint8
	rangingData        vl53l1xRangingData
	results            vl53l1xResultBuffer
}

//-------------------------------------------------------------------------------------
// VL53L1x
func newVl53l1x(bus drivers.I2C, addr uint16) interface{} {
	if addr != 0 {
		return Vl53l1x{
			bus:     bus,
			Address: addr,
			mode:    VL53L1x_LONG,
			timeout: 500,
		}
	} else {
		return Vl53l1x{
			bus:     bus,
			Address: Vl53l1xAddress,
			mode:    VL53L1x_LONG,
			timeout: 500,
		}
	}
}

//-------------------------------------------------------------------------------------
// VL53L1x
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
