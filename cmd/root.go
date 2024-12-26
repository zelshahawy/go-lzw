package cmd

import (
	"fmt"
)

func StartEncoding(fileName string) {
	fmt.Printf("Encoding has started on %v", fileName)
}

func StartDecoding(fileName string) {
	fmt.Printf("Decoding has started on %v", fileName)
}
