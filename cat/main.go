package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	// No arguments: read from stdin
	if len(args) == 0 { // No arguments provided
		_, err := io.Copy(os.Stdout, os.Stdin) // Read from stdin and write to stdout
		if err != nil {                        // Handle error
			fmt.Fprintln(os.Stderr, "ERROR:", err) // Print error message to stderr
			os.Exit(1)                             // Exit with error code
		}
		return // Exit after copying stdin to stdout
	}

	// Arguments provided: treat as filenames
	for _, filename := range args { // Iterate over each filename
		file, err := os.Open(filename) // Open the file
		if err != nil {                // Handle error
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err) // Print error message to stderr
			os.Exit(1)                                 // Exit with error code
		}
		io.Copy(os.Stdout, file) // Copy file content to stdout
		file.Close()             // Close the file
	}
}
