package internal

import (
	"fmt"

	"github.com/zelshahawy/go-lzw/internal/dictionary"
)

func EncodeFile() {
	fmt.Println("Encoding has started")
	dict, lookup := dictionary.InitDictionary()
	fmt.Println(dict)
	fmt.Println(lookup)
}

func DecodeFile() {
	fmt.Println("Decoding has started")
	dict, lookup := dictionary.InitDictionary()
	fmt.Println(dict)
	fmt.Println(lookup)
}
