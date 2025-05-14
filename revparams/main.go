package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Loop backwards from last arg to the first (excluding program name)
	for i := len(os.Args) - 1; i > 0; i-- { // Iterate over each argument in reverse order
		for _, r := range os.Args[i] { // Iterate over each character in the argument
			z01.PrintRune(r) //	 Print the character
		}
		z01.PrintRune('\n') // Print a newline after each argument
	}
}
