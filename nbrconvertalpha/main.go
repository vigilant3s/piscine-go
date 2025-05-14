package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// Helper function to print a single rune
func printRune(r rune) {
	z01.PrintRune(r)
}

// Helper function to print a string rune-by-rune
func printStr(s string) {
	for _, r := range s {
		printRune(r)
	}
}

func main() {
	args := os.Args[1:]
	upper := false

	// Check for --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:] // remove the flag from args
	}

	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil || n < 1 || n > 26 {
			printRune(' ')
			continue
		}
		// Convert number to letter (ASCII logic)
		if upper {
			printRune(rune('A' + n - 1))
		} else {
			printRune(rune('a' + n - 1))
		}
	}
	// End with newline for cleaner output
	printRune('\n')
}
