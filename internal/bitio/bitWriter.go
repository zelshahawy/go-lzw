package bitio

func (bp *BitPacker) WriteCode(code int, codeSize int) {
	// Ensure the code is padded to the left to match the codeSize
	code &= (1 << codeSize) - 1

	bp.bitBuf |= (uint64(code) << bp.bitCount)
	bp.bitCount += codeSize

	// Flush any full bytes (8 bits) from bitBuf
	for bp.bitCount >= 8 {
		b := byte(bp.bitBuf & 0xFF)
		bp.output = append(bp.output, b)
		bp.bitBuf >>= 8
		bp.bitCount -= 8
	}
}

// FlushRemaining writes out any leftover bits (less than a byte).
func (bp *BitPacker) FlushRemaining() {
	if bp.bitCount > 0 {
		// We still have a partial byte in bitBuf
		b := byte(bp.bitBuf & 0xFF)
		bp.output = append(bp.output, b)
		// Reset buffer
		bp.bitBuf = 0
		bp.bitCount = 0
	}
}
