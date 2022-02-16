package multi

import (
	"time"
)

type Amg88xxInterruptMode uint8

type Amg88xx struct {
	bus             I2C
	Address         uint16
	data            []uint8
	interruptMode   Amg88xxInterruptMode
	interruptEnable uint8
}

func newAmg88xx(bus I2C, addr uint16) Amg88xx {
	if addr != 0 {
		return Amg88xx{
			bus:     bus,
			Address: addr,
		}
	} else {
		return Amg88xx{
			bus:     bus,
			Address: Amg88xxAddressHigh,
		}
	}
}

func (d Amg88xx) configure() error {
	d.data = make([]uint8, 128)

	d.SetPCTL(AMG88XX_NORMAL_MODE)
	d.SetReset(AMG88XX_INITIAL_RESET)
	d.SetFrameRate(AMG88XX_FPS_10)

	time.Sleep(100 * time.Millisecond)

	return nil
}

// ReadPixels returns the 64 values (8x8 grid) of the sensor converted to  millicelsius
func (d Amg88xx) ReadPixels(buffer *[64]int16) {
	d.bus.ReadRegister(uint8(d.Address), AMG88XX_PIXEL_OFFSET, d.data)
	for i := 0; i < 64; i++ {
		buffer[i] = int16((uint16(d.data[2*i+1]) << 8) | uint16(d.data[2*i]))
		if (buffer[i] & (1 << 11)) > 0 { // temperature negative
			buffer[i] &= ^(1 << 11)
			buffer[i] = -buffer[i]
		}
		buffer[i] *= AMG88XX_PIXEL_TEMP_CONVERSION
	}
}

func (d Amg88xx) SetPCTL(pctl uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_PCTL, []byte{pctl})
}

func (d Amg88xx) SetReset(rst uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_RST, []byte{rst})
}

func (d Amg88xx) SetFrameRate(framerate uint8) {
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_FPSC, []byte{framerate & 0x01})
}

func (d Amg88xx) SetMovingAverageMode(mode bool) {
	var value uint8
	if mode {
		value = 1
	}
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_AVE, []byte{value << 5})
}

func (d Amg88xx) SetInterruptLevels(high int16, low int16) {
	d.SetInterruptLevelsHysteresis(high, low, (high*95)/100)
}

func (d Amg88xx) SetInterruptLevelsHysteresis(high int16, low int16, hysteresis int16) {
	high = high / AMG88XX_PIXEL_TEMP_CONVERSION
	if high < -4095 {
		high = -4095
	}
	if high > 4095 {
		high = 4095
	}
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8(high & 0xFF)})
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8((high & 0xFF) >> 4)})

	low = low / AMG88XX_PIXEL_TEMP_CONVERSION
	if low < -4095 {
		low = -4095
	}
	if low > 4095 {
		low = 4095
	}
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8(low & 0xFF)})
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8((low & 0xFF) >> 4)})

	hysteresis = hysteresis / AMG88XX_PIXEL_TEMP_CONVERSION
	if hysteresis < -4095 {
		hysteresis = -4095
	}
	if hysteresis > 4095 {
		hysteresis = 4095
	}
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8(hysteresis & 0xFF)})
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTHL, []byte{uint8((hysteresis & 0xFF) >> 4)})
}

func (d Amg88xx) EnableInterrupt() {
	d.interruptEnable = 1
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTC, []byte{((uint8(d.interruptMode) << 1) | d.interruptEnable) & 0x03})
}

func (d Amg88xx) DisableInterrupt() {
	d.interruptEnable = 0
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTC, []byte{((uint8(d.interruptMode) << 1) | d.interruptEnable) & 0x03})
}

func (d Amg88xx) SetInterruptMode(mode Amg88xxInterruptMode) {
	d.interruptMode = mode
	d.bus.WriteRegister(uint8(d.Address), AMG88XX_INTC, []byte{((uint8(d.interruptMode) << 1) | d.interruptEnable) & 0x03})
}

func (d Amg88xx) GetInterrupt() []uint8 {
	data := make([]uint8, 8)
	d.bus.ReadRegister(uint8(d.Address), AMG88XX_INT_OFFSET, data)
	return data
}

func (d Amg88xx) ClearInterrupt() {
	d.SetReset(AMG88XX_FLAG_RESET)
}

func (d Amg88xx) ReadThermistor() int16 {
	data := make([]uint8, 2)
	d.bus.ReadRegister(uint8(d.Address), AMG88XX_TTHL, data)
	return (int16((uint16(data[1])<<8)|uint16(data[0])) * AMG88XX_THERMISTOR_CONVERSION) / 10
}
