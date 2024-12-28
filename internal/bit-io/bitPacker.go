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

// Bytes returns the packed data as a byte slice.
func (bp *BitPacker) Bytes() []byte {
	return bp.output
}
