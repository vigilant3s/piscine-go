package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Custom boolean type
type boolean int

// Boolean constants
const (
	no  boolean = 0
	yes boolean = 1
)

// Messages
const (
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)

// Print string with newline using z01
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Helper function: 1 if even, 0 if odd
func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

// Convert int to custom boolean
func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	}
	return no
}

func main() {
	lengthOfArg := len(os.Args[1:]) // Skip program name
	if isEven(lengthOfArg) == yes {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
