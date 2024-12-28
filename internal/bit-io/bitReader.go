package bitio

import (
	"errors"
)

var ErrUnexpectedEOF = errors.New("unexpected EOF")

// BitPacker is a struct that holds the packed data.

func (bp *BitPacker) ReadCode(codeSize int) (int, error) {
	code := 0
	for i := 0; i < codeSize; i++ {
		if bp.bitCount == 0 {
			return 0, ErrUnexpectedEOF
		}

		// Get the next bit
		bit := (bp.bitBuf & 1)
		bp.bitBuf >>= 1
		bp.bitCount--

		// Add the bit to the code
		code |= int(bit) << i
	}

	return code, nil
}
