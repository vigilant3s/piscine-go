package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Atoi manually converts a string to an integer.
// Returns -1 if the string contains any non-digit characters.
func Atoi(s string) int {
	result := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1 // invalid input
		}
		result = result*10 + int(ch-'0')
	}
	return result
}

// PrintString prints a string using z01.PrintRune
func PrintString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	args := os.Args[1:]

	uppercase := false
	startIndex := 0

	// Check for --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		uppercase = true
		startIndex = 1
	}

	for i := startIndex; i < len(args); i++ {
		n := Atoi(args[i])
		if n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		if uppercase {
			z01.PrintRune(rune('A' + n - 1))
		} else {
			z01.PrintRune(rune('a' + n - 1))
		}
	}
	z01.PrintRune('\n')
}
