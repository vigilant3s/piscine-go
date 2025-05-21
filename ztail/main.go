package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Must have at least "-c" and a count plus one file
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		fmt.Println("Usage: go run . -c <count> <file1> [file2] [...]")
		os.Exit(1)
	}

	// Parse the count after "-c"
	count, err := strconv.Atoi(os.Args[2])
	if err != nil || count < 0 {
		fmt.Println("Invalid count:", os.Args[2])
		os.Exit(1)
	}

	files := os.Args[3:]    // All remaining args are file names
	multi := len(files) > 1 // Whether we have multiple files
	hadError := false       // Track if any file produced an error

	for i, name := range files {
		// Print header for multiple files
		if multi {
			if i > 0 {
				fmt.Println() // Blank line between sections
			}
			fmt.Printf("==> %s <==\n", name)
		}

		// Read entire file into memory
		data, err := os.ReadFile(name)
		if err != nil {
			// Print the raw error message
			fmt.Println(err.Error())
			hadError = true
			continue
		}

		// Determine start index for last count bytes
		start := len(data) - count
		if start < 0 {
			start = 0
		}

		// Print only the tail slice
		fmt.Printf("%s", data[start:])
	}

	// Exit with non-zero status if any errors occurred
	if hadError {
		os.Exit(1)
	}
}
