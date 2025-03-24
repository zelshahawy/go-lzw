package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/zelshahawy/go-lzw/cmd"
)

func main() {
	// Parse command name
	cmdName := filepath.Base(os.Args[0])

	if len(os.Args) > 9 {
		fmt.Fprintln(os.Stderr, "Too many filenames provided, maximum is 8")
		os.Exit(1)
	}

	var wg sync.WaitGroup // WaitGroup for synchronization
	if len(os.Args) < 2 {
		// No filenames provided, read from stdin
		wg.Add(1)
		go func() {
			defer wg.Done()
			switch cmdName {
			case "encode":
				fmt.Print("Enter text to encode (Ctrl+D to finish): ")
				cmd.StartEncoding(os.Stdin, "stdin") // Use EOF to signal end of input. Ctrl+D on Unix systems.
			case "decode":
				fmt.Print("Enter text to decode (Ctrl+D to finish): ")
				cmd.StartDecoding(os.Stdin, "stdin")
			default:
				fmt.Fprintln(os.Stderr, "Invalid command name")
				os.Exit(1)
			}
		}()
	} else {
		// Filenames provided
		for i := 1; i < len(os.Args); i++ {
			filename := os.Args[i]
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()
				file, err := os.Open(filename)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error opening file: %s\n", filename)
					return
				}
				defer file.Close()

				switch cmdName {
				case "encode":
					cmd.StartEncoding(file, filepath.Base(filename))
				case "decode":
					cmd.StartDecoding(file, filepath.Base(filename))
				default:
					fmt.Fprintln(os.Stderr, "Invalid command name")
					os.Exit(1)
				}
			}(filename)
		}
	}

	wg.Wait()
	// fmt.Println("All goroutines have finished.")
}
