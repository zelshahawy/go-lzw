package internal

import (
	"io"
	"os"

	"github.com/zelshahawy/go-lzw/internal/bitio"
	"github.com/zelshahawy/go-lzw/internal/dictionary"
)

// reconstructString traverses the chain (Prefix, Ch) until we reach -1 (meaning a
// single-byte code). It collects the bytes in reverse, then reverses them to
// produce the full string.
func reconstructString(dict []dictionary.DictionaryEntry, code int) []byte {
	if code < 256 {
		// Single-byte character
		return []byte{byte(code)}
	}

	var result []byte
	for code != -1 {
		entry := dict[code]
		result = append(result, entry.Ch)
		code = entry.Prefix
	}
	// Reverse in-place
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// ExecDecoding reads all compressed data from 'input', decodes it via LZW,
// accumulates the decoded data in memory, then writes it once at the end
// to minimize I/O.
func ExecDecoding(input io.Reader) error {
	// 1) Prepare a BitPacker for decoding (it has Output, CodesOutput, BitBuf, BitCount).
	bp := bitio.NewBitPacker()

	// 2) Initialize the LZW dictionary: codes 0..255 are single-byte entries.
	dict, _ := dictionary.InitDictionary()
	nextCode := 256
	codeSize := 9

	// We'll track the "oldCode" and "oldString" as we decode each code.
	oldCode := -1
	var oldString []byte

	// 3) Define a helper function that decodes a single code, returning the
	//    uncompressed bytes. We keep it inside ExecDecoding so it can
	//    capture dict, nextCode, oldCode, oldString, and codeSize by reference.
	decodeOneCode := func(newCode int) []byte {
		// If this is the very first code in the stream:
		if oldCode == -1 {
			oldCode = newCode
			oldString = reconstructString(dict, newCode)
			return oldString
		}

		// Normal LZW decode flow, including the “K-W” edge case:
		var newString []byte
		if newCode >= len(dict) && newCode == nextCode {
			// K-W: newString = oldString + firstByte(oldString)
			firstByte := oldString[0]
			newString = append(oldString, firstByte)
		} else {
			// Else reconstruct from the dictionary
			newString = reconstructString(dict, newCode)
		}

		// Add a new entry to the dictionary: oldString + first byte of newString
		dict = append(dict, dictionary.DictionaryEntry{
			Prefix: oldCode,
			Ch:     newString[0],
		})
		nextCode++

		// Possibly increase codeSize if we've used all codes at this size
		if nextCode == (1<<codeSize) && codeSize < 15 {
			codeSize++
		}

		// Update oldCode / oldString
		oldCode = newCode
		oldString = newString
		return newString
	}

	// 4) Read compressed bytes from `input`, feed them into bp.BitBuf,
	//    extract codes, decode them, and accumulate the result in memory.
	buf := make([]byte, 4096)
	for {
		n, readErr := input.Read(buf)
		if n > 0 {
			for i := 0; i < n; i++ {
				b := buf[i]
				// Shift this byte into bitBuf
				bp.BitBuf |= (uint64(b) << bp.BitCount)
				bp.BitCount += 8

				// Extract as many `codeSize`-bit codes as possible
				for {
					// log.Printf("bp.BitCount: %d\n", bp.BitCount)
					if bp.BitCount < codeSize {
						break
					}
					code := bp.ReadCode(codeSize)

					if code == -1 {
						// log.Printf("code == -1\n")
						break
					}
					// Optionally store the code in bp.CodesOutput
					bp.CodesOutput = append(bp.CodesOutput, code)

					// Decode the code => uncompressed bytes
					decodedBytes := decodeOneCode(code)

					// Accumulate decoded bytes in bp.Output
					bp.Output = append(bp.Output, decodedBytes...)
				}
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return readErr
		}
	}

	// 5) After EOF, there might still be leftover bits for one or more codes.
	for {
		code := bp.ReadCode(codeSize)
		if code == -1 {
			break
		}
		bp.CodesOutput = append(bp.CodesOutput, code)

		decodedBytes := decodeOneCode(code)
		bp.Output = append(bp.Output, decodedBytes...)
	}

	// 6) Now bp.Output contains all uncompressed data in memory.
	//    We'll do one final write to either stdout (if CLI=1) or output.out.

	if getEnv("CLI") == "1" {
		_, err := os.Stdout.Write(bp.Output)
		return err
	} else {
		outF, err := os.Create("output.out")
		if err != nil {
			return err
		}
		defer outF.Close()

		_, err = outF.Write(bp.Output)
		return err
	}
}
