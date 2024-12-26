package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zelshahawy/go-lzw/cmd"
)

func main() {
	fmt.Println("Entry Point called")
	cmdName := filepath.Base(os.Args[0])

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  encode <filename>")
		fmt.Println("  decode <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]

	switch cmdName {
	case "encode":
		cmd.StartEncoding(filename)

	case "decode":
		cmd.StartDecoding(filename)

	default:
		fmt.Println("Wrong Command Name Given")
	}
}
