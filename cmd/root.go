package cmd

import (
	"fmt"
	"io"

	"github.com/zelshahawy/go-lzw/internal"
)

// StartEncoding function
// It starts the encoding process
// it accepts an io.Reader and a filename
// It calls the internal.ExecEncoding function
func StartEncoding(input io.Reader, filename string) {
	if err := internal.ExecEncoding(input, filename); err != nil {
		fmt.Println("Error Encoding")
	}
}

// StartDecoding function
// It starts the decoding process
// it accepts an io.Reader and a filename
// It calls the internal.ExecDecoding function
func StartDecoding(input io.Reader, filename string) {
	if err := internal.ExecDecoding(input, filename); err != nil {
		fmt.Println("Error Decoding")
	}
}
