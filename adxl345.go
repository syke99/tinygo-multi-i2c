package multi

type Adxl345Range uint8
type Adxl345Rate uint8

// Internal structure for the power configuration
type adxl345PowerCtl struct {
	link      uint8
	autoSleep uint8
	measure   uint8
	sleep     uint8
	wakeUp    uint8
}

// Internal structure for the sensor's data format configuration
type adxl345DataFormat struct {
	selfTest    uint8
	spi         uint8
	intInvert   uint8
	fullRes     uint8
	justify     uint8
	sensorRange Adxl345Range
}

// Internal structure for the sampling rate configuration
type adxl345BwRate struct {
	lowPower uint8
	rate     Adxl345Rate
}

type Adxl345 struct {
	bus        I2C
	Address    uint16
	powerCtl   adxl345PowerCtl
	dataFormat adxl345DataFormat
	bwRate     adxl345BwRate
}

func newAdx1345(bus I2C, addr uint16) Adxl345 {
	if addr != 0 {
		return Adxl345{
			bus: bus,
			powerCtl: adxl345PowerCtl{
				measure: 1,
			},
			dataFormat: adxl345DataFormat{
				sensorRange: ADX1345_RANGE_2G,
			},
			bwRate: adxl345BwRate{
				lowPower: 1,
				rate:     ADX1345_RATE_100HZ,
			},
			Address: addr,
		}
	} else {
		return Adxl345{
			bus: bus,
			powerCtl: adxl345PowerCtl{
				measure: 1,
			},
			dataFormat: adxl345DataFormat{
				sensorRange: ADX1345_RANGE_2G,
			},
			bwRate: adxl345BwRate{
				lowPower: 1,
				rate:     ADX1345_RATE_100HZ,
			},
			Address: Adx1345AddressLow,
		}
	}
}

func (d Adxl345) configure() error {
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_BW_RATE, []byte{d.bwRate.toByte()})
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_POWER_CTL, []byte{d.powerCtl.toByte()})
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_DATA_FORMAT, []byte{d.dataFormat.toByte()})

	return nil
}

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

// Halt stops the sensor, values will not updated
func (d Adxl345) Halt() {
	d.powerCtl.measure = 0
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_POWER_CTL, []byte{d.powerCtl.toByte()})
}

// Restart makes reading the sensor working again after a halt
func (d Adxl345) Restart() {
	d.powerCtl.measure = 1
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_POWER_CTL, []byte{d.powerCtl.toByte()})
}

// ReadAcceleration reads the current acceleration from the device and returns
// it in µg (micro-gravity). When one of the axes is pointing straight to Earth
// and the sensor is not moving the returned value will be around 1000000 or
// -1000000.
func (d Adxl345) ReadAcceleration() (x int32, y int32, z int32, err error) {
	rx, ry, rz := d.ReadRawAcceleration()

	x = d.dataFormat.convertToIS(rx)
	y = d.dataFormat.convertToIS(ry)
	z = d.dataFormat.convertToIS(rz)

	return
}

// ReadRawAcceleration reads the sensor values and returns the raw x, y and z axis
// from the adxl345.
func (d Adxl345) ReadRawAcceleration() (x int32, y int32, z int32) {
	data := []byte{0, 0, 0, 0, 0, 0}
	d.bus.ReadRegister(uint8(d.Address), ADX1345_REG_DATAX0, data)

	x = readAdxl345IntLE(data[0], data[1])
	y = readAdxl345IntLE(data[2], data[3])
	z = readAdxl345IntLE(data[4], data[5])

	return
}

// UseLowPower sets the ADXL345 to use the low power mode.
func (d Adxl345) UseLowPower(power bool) {
	if power {
		d.bwRate.lowPower = 1
	} else {
		d.bwRate.lowPower = 0
	}
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_BW_RATE, []byte{d.bwRate.toByte()})
}

// SetRate change the current rate of the sensor
func (d Adxl345) SetRate(rate Adxl345Rate) bool {
	d.bwRate.rate = rate & 0x0F
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_BW_RATE, []byte{d.bwRate.toByte()})
	return true
}

// SetRange change the current range of the sensor
func (d Adxl345) SetRange(sensorRange Adxl345Range) bool {
	d.dataFormat.sensorRange = sensorRange & 0x03
	d.bus.WriteRegister(uint8(d.Address), ADX1345_REG_DATA_FORMAT, []byte{d.dataFormat.toByte()})
	return true
}

// private functions

// convertToIS adjusts the raw values from the adxl345 with the range configuration
func (d adxl345DataFormat) convertToIS(rawValue int32) int32 {
	switch d.sensorRange {
	case ADX1345_RANGE_2G:
		return rawValue * 4 // rawValue * 2 * 1000 / 512
	case ADX1345_RANGE_4G:
		return rawValue * 8 // rawValue * 4 * 1000 / 512
	case ADX1345_RANGE_8G:
		return rawValue * 16 // rawValue * 8 * 1000 / 512
	case ADX1345_RANGE_16G:
		return rawValue * 32 // rawValue * 16 * 1000 / 512
	default:
		return 0
	}
}

// readInt converts two bytes to int16
func readAdxl345IntLE(msb byte, lsb byte) int32 {
	return int32(uint16(msb) | uint16(lsb)<<8)
}
