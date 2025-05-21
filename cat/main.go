package main

import (
	"io"
	"os"
)

func main() { // cat.go
	args := os.Args[1:] // Ignore the first argument (the program name)

	// No arguments → read from stdin
	if len(args) == 0 { // cat stdin
		_, err := io.Copy(os.Stdout, os.Stdin) // Copy from stdin to stdout
		if err != nil {                        // Handle error
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n") // Write error to stderr
			os.Exit(1)                                            // Exit with error code
		}
		return // Exit if no arguments
	}

	// Loop over each argument
	for _, filename := range args { // cat file1 file2 ...
		file, err := os.Open(filename) // Open the file
		if err != nil {                // Handle error
			os.Stderr.WriteString("ERROR: open " + filename + ": " + err.Error() + "\n") // Write error to stderr
			os.Exit(1)                                                                   // Exit with error code
		}
		io.Copy(os.Stdout, file) // Copy from file to stdout
		file.Close()             // Close the file
	}
}
