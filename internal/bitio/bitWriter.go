package bitio

import "os"

func (bp *BitPacker) WriteCode(code int, codeSize int) {
	// Ensure the code is padded to the left to match the codeSize
	code &= (1 << codeSize) - 1

	bp.BitBuf |= (uint64(code) << bp.BitCount)
	bp.BitCount += codeSize

	// Flush any full bytes (8 bits) from bitBuf
	for bp.BitCount >= 8 {
		b := byte(bp.BitBuf & 0xFF)
		bp.Output = append(bp.Output, b)
		bp.BitBuf >>= 8
		bp.BitCount -= 8
	}
}

// FlushRemaining writes out any leftover bits (less than a byte).
func (bp *BitPacker) FlushRemaining() {
	if bp.BitCount > 0 {
		// We still have a partial byte in bitBuf
		b := byte(bp.BitBuf & 0xFF)
		bp.Output = append(bp.Output, b)
		// Reset buffer
		bp.BitBuf = 0
		bp.BitCount = 0
	}
}

func (bp *BitPacker) WriteOutputToFile() error {
	outFile := "output.lzw"
	if getEnv("CLI") == "1" {
		_, err := os.Stdout.Write(bp.Bytes())
		if err != nil {
			return err
		}
	} else {
		outF, err := os.Create(outFile)
		if err != nil {
			return err
		}
		defer outF.Close()

		_, err = outF.Write(bp.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}
