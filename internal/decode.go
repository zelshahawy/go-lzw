package internal

import (
	"fmt"
	"io"
	"os"

	"github.com/zelshahawy/go-lzw/internal/bitio"
	"github.com/zelshahawy/go-lzw/internal/dictionary"
)

func ExecDecoding(fileName string) error {
	fmt.Println("Decoding has started")
	dict, lookup := dictionary.InitDictionary()
	nextCode := 256
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error Opening File")
		return err
	}
	defer file.Close()

	bp := bitio.NewBitPacker()
	codeSize := 9

	// p = current prefix code
	p := -1

	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
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
				}

				// Add new entry to dictionary
				dict = append(dict, dictionary.DictionaryEntry{Prefix: p, Ch: c})
				lookup[pc] = nextCode
				nextCode++

				// Possibly increase codeSize if we hit the limit
				if nextCode == (1<<codeSize) && codeSize < 15 {
					codeSize++
				}

				// Reset prefix
				p = int(c)
			}
		}
	}

	// Output last code
	if p != -1 {
		bp.WriteCode(p, codeSize)
	}
	return nil
}
