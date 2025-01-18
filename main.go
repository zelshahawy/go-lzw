package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/zelshahawy/go-lzw/cmd"
)

// Starter function for the program
// It reads the command name and the filenames from the command line arguments
// It checks if the command name is either "encode" or "decode" and the number of filenames is valid
// Then it opens the files and starts the encoding/decoding process
// It uses a WaitGroup to synchronize the goroutines
// It waits for all goroutines to finish before exiting
func main() {
	// log.Print("Entry Point called")
	cmdName := filepath.Base(os.Args[0])

	if len(os.Args) < 2 {
		fmt.Println("No filenames provided")
		os.Exit(1)
	}

	if len(os.Args) > 9 {
		fmt.Println("Too many filenames provided, maximum is 8")
		os.Exit(1)
	}

	var inputs []io.Reader
	var fileNames []string
	var wg sync.WaitGroup // WaitGroup for synchronization

	for i := 1; i < len(os.Args); i++ {
		filename := os.Args[i]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error Opening File:", filename)
			os.Exit(1)
		}
		defer file.Close()
		inputs = append(inputs, file)
		fileNames = append(fileNames, filepath.Base(filename))
	}

	switch cmdName {
	case "encode":
		for i, input := range inputs {
			wg.Add(1)
			go func(input io.Reader, filename string) {
				defer wg.Done()
				cmd.StartEncoding(input, filename)
			}(input, fileNames[i])
		}

	case "decode":
		for i, input := range inputs {
			wg.Add(1)
			go func(input io.Reader, filename string) {
				defer wg.Done()
				cmd.StartDecoding(input, filename)
			}(input, fileNames[i])
		}

	default:
		fmt.Println("Wrong Command Name Given")
	}
	wg.Wait()
}
