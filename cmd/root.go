package cmd

import (
	"fmt"

	"github.com/zelshahawy/go-lzw/internal"
)

func StartEncoding(fileName string) {
	fmt.Printf("Encoding has started on %v\n", fileName)
	if internal.ExecEncoding(fileName) != nil {
		fmt.Println("Error Encoding")
	}
}

func StartDecoding(fileName string) {
	fmt.Printf("Decoding has started on %v\n", fileName)
	if internal.ExecDecoding(fileName) != nil {
		fmt.Println("Error Decoding")
	}
}
