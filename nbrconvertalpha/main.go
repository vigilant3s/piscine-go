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
			return -1
		}
		result = result*10 + int(ch-'0')
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return // nothing to print, not even a newline
	}

	uppercase := false
	start := 0

	if args[0] == "--upper" {
		uppercase = true
		start = 1
	}

	for i := start; i < len(args); i++ {
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
}
