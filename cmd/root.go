package cmd

import (
	"fmt"
	"io"

	"github.com/zelshahawy/go-lzw/internal"
)

func StartEncoding(input io.Reader) {
	if err := internal.ExecEncoding(input); err != nil {
		fmt.Println("Error Encoding")
	}
}

func StartDecoding(input io.Reader) {
	if err := internal.ExecDecoding(input); err != nil {
		fmt.Println("Error Decoding")
	}
}
