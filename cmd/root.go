package cmd

import (
	"fmt"
)

func StartEncoding(fileName string) {
	fmt.Printf("Encoding has started on %v\n", fileName)
}

func StartDecoding(fileName string) {
	fmt.Printf("Decoding has started on %v\n", fileName)
}
