package cmd

import (
	"fmt"

	"github.com/zelshahawy/go-lzw/internal"
)

func StartEncoding(fileName string) {
	fmt.Printf("Encoding has started on %v\n", fileName)
	internal.EncodeFile(fileName)
}

func StartDecoding(fileName string) {
	fmt.Printf("Decoding has started on %v\n", fileName)
	internal.DecodeFile(fileName)
}
