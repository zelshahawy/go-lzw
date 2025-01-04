package cmd

import (
	"fmt"
	"io"

	"github.com/zelshahawy/go-lzw/internal"
)

func StartEncoding(input io.Reader, filename string) {
	if err := internal.ExecEncoding(input, filename); err != nil {
		fmt.Println("Error Encoding")
	}
}

func StartDecoding(input io.Reader, filename string) {
	if err := internal.ExecDecoding(input, filename); err != nil {
		fmt.Println("Error Decoding")
	}
}
