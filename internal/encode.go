package internal

import (
	"fmt"
	"io"
	"os"

	"github.com/zelshahawy/go-lzw/internal/bitio"
	"github.com/zelshahawy/go-lzw/internal/dictionary"
)

func ExecEncoding(fileName string) error {
	fmt.Println("Encoding has started")
	dict, lookup := dictionary.InitDictionary()
	nextCode := 256

	outFile := fileName + ".lzw"
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

	// Write the packed bytes to outFile
	outF, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer outF.Close()

	_, err = outF.Write(bp.Bytes())
	return err
}
