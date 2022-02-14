package multi

//-------------------------------------------------------------------------------------
// Universal configuration methods
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
