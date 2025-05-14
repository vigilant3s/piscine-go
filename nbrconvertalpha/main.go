package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Atoi converts a string to an integer.
// Returns -1 if the string contains non-digit characters.
func Atoi(s string) int {
	result := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1
		}
		result = result*10 + int(ch-'0')
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return // Nothing to print, not even a newline
	}

	uppercase := false
	start := 0

	if args[0] == "--upper" {
		uppercase = true
		start = 1
	}

	printed := false

	for i := start; i < len(args); i++ {
		n := Atoi(args[i])
		if n < 1 || n > 26 {
			z01.PrintRune(' ')
		} else {
			if uppercase {
				z01.PrintRune(rune('A' + n - 1))
			} else {
				z01.PrintRune(rune('a' + n - 1))
			}
		}
		printed = true
	}

	// Print newline only if we actually printed something
	if printed {
		z01.PrintRune('\n')
	}
}
