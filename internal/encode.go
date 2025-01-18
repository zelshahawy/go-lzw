package internal

import (
	"io"

	"github.com/zelshahawy/go-lzw/internal/bitio"
	"github.com/zelshahawy/go-lzw/internal/dictionary"
)

// ExecEncoding reads all data from 'input', encodes it via LZW,
// accumulates the encoded data in memory, then writes it once at the end to minimize I/O.
// Accepts an io.Reader and a filename

func ExecEncoding(input io.Reader, filename string) error {
	// log.Printf("Encoding has started\n##############################\n\n")
	dict, lookup := dictionary.InitDictionary()
	nextCode := 256

	bp := bitio.NewBitPacker()
	codeSize := 9
	const MAXBITS int = 20
	const MAXCODE int = 1<<MAXBITS - 1

	// p = current prefix code
	p := -1

	buf := make([]byte, 4096)
	for {
		n, err := input.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			// No more data
			break
		}

		for _, c := range buf[:n] {
			pc := [2]int{p, int(c)}
			if code, found := lookup[pc]; found {
				// If (p, c) is known, extend prefix
				p = code
			} else {
				// Output current prefix
				if p != -1 {
					// Write p using current codeSize
					bp.WriteCode(p, codeSize)
					// log.Printf("ENC: code=%d, codeSize=%d, nextCode=%d, p=%d, c=%d", p, codeSize, nextCode, p, c)

				}

				// Add new entry to dictionary

				if nextCode < MAXCODE {
					dict = append(dict, dictionary.DictionaryEntry{Prefix: p, Ch: c})
					lookup[pc] = nextCode
					nextCode++
					if nextCode == (1<<codeSize) && codeSize < MAXBITS {
						codeSize++
					}
				}

				// Reset p to the code for the single char c
				p = int(c)
			}
		}

		if err == io.EOF {
			break
		}
	}

	// Output the final prefix
	if p != -1 {
		bp.WriteCode(p, codeSize)
	}

	// Flush leftover bits
	bp.FlushRemaining()

	// Write the packed bytes to outFile or stdout
	fileName := (filename) + ".lzw"
	bp.WriteOutputToFile(fileName)
	return nil
}
