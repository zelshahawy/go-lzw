package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/zelshahawy/go-lzw/cmd"
)

func main() {
	// log.Print("Entrty Point called")
	cmdName := filepath.Base(os.Args[0])

	var filename string
	var input io.Reader

	if len(os.Args) < 2 {
		// log.Print("No filename provided, reading from stdin")
		input = os.Stdin
	} else {
		filename = os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error Opening File")
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	switch cmdName {
	case "encode":
		cmd.StartEncoding(input)

	case "decode":
		cmd.StartDecoding(input)

	default:
		fmt.Println("Wrong Command Name Given")
	}
}
