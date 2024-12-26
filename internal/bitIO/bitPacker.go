package bitio

type BitPacker struct {
	output   []byte // where we store packed bytes
	bitBuf   uint64 // buffer for bits not yet written
	bitCount int    // how many bits are currently in bitBuf
}

// NewBitPacker creates an empty BitPacker.
func NewBitPacker() *BitPacker {
	return &BitPacker{
		output:   make([]byte, 0),
		bitBuf:   0,
		bitCount: 0,
	}
}

// WriteCode writes `code` using `codeSize` bits into the buffer.
func (bp *BitPacker) WriteCode(code int, codeSize int) {
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

// Bytes returns the packed data as a byte slice.
func (bp *BitPacker) Bytes() []byte {
	return bp.output
}
