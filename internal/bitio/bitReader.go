package bitio

// BitPacker is a struct that holds the packed data.

func (bp *BitPacker) ReadCode(codeSize int) int {
	code := 0
	for i := 0; i < codeSize; i++ {
		if bp.BitCount == 0 {
			return -1
		}

		// Get the next bit
		bit := (bp.BitBuf & 1)
		bp.BitBuf >>= 1
		bp.BitCount--

		// Add the bit to the code
		code |= int(bit) << i
	}

	return code
}
