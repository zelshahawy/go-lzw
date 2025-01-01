package bitio

type BitPacker struct {
	Output   []byte // where we store packed bytes
	BitBuf   uint64 // buffer for bits not yet written
	BitCount int    // how many bits are currently in bitBuf
}

// NewBitPacker creates an empty BitPacker.
func NewBitPacker() *BitPacker {
	return &BitPacker{
		Output:   make([]byte, 0),
		BitBuf:   0,
		BitCount: 0,
	}
}

// Bytes returns the packed data as a byte slice.
func (bp *BitPacker) Bytes() []byte {
	return bp.Output
}
