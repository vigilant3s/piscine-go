package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Start from index 1 to skip the program name
	for i := 1; i < len(os.Args); i++ { // Iterate over each argument
		for _, r := range os.Args[i] { // Iterate over each character in the argument
			z01.PrintRune(r) // Print the character
		}
		// After each argument, print a newline
		z01.PrintRune('\n')
	}
}
