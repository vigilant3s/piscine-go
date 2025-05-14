package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() { // Print the program name
	programPath := os.Args[0] // Get the program path
	startIndex := 0           // Initialize startIndex to 0

	// Find the last slash in the path
	for i, char := range programPath { // Iterate over each character
		if char == '/' { // Check if the character is a slash
			startIndex = i + 1 // Update startIndex to the character after the slash
		}
	}

	// Print from the character after the last slash
	for _, char := range programPath[startIndex:] { // Iterate over each character starting from startIndex
		z01.PrintRune(char) // Print the character
	}

	z01.PrintRune('\n') // Print a newline character
}
