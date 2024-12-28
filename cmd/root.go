package cmd

import (
	"fmt"

	"github.com/zelshahawy/go-lzw/internal"
)

func StartEncoding(fileName string) {
	fmt.Printf("Encoding has started on %v\n", fileName)
	if err := internal.ExecEncoding(fileName); err != nil {
		fmt.Println("Error Encoding")
	}
}

func StartDecoding(fileName string) {
	fmt.Printf("Decoding has started on %v\n", fileName)
	if err := internal.ExecDecoding(fileName); err != nil {
		fmt.Println("Error Decoding")
	}
}
