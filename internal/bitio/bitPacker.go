package bitio

// BitPacker is a helper struct for packing bits into bytes.
// It is used by the encoder to write the packed data to the output file.
// It has three fields: Output, BitBuf, and BitCount.

type BitPacker struct {
	Output   []byte // where we store packed bytes
	BitBuf   uint64 // buffer for bits not yet written
	BitCount int    // how many bits are currently in bitBuf
}

// WriteBits writes the given value to the output buffer.
// It writes the value to the buffer, then shifts the buffer to the left by the number of bits written.
func NewBitPacker() *BitPacker {
	return &BitPacker{
		Output:   make([]byte, 0),
		BitBuf:   0,
		BitCount: 0,
	}
}

// WriteBits writes the given value to the output buffer.
// Bytes returns the packed data as a byte slice.
func (bp *BitPacker) Bytes() []byte {
	return bp.Output
}
