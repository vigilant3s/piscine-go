package piscine

import "github.com/01-edu/z01"

// PrintNbrBase prints an integer in a custom base
func PrintNbrBase(nbr int, base string) {
	// Check if base is valid
	if !isValidBase(base) { //
		printStr("NV")
		return
	}

	if nbr == 0 { //
		z01.PrintRune(rune(base[0]))
		return
	}

	// Handle negative
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	// Convert to base
	baseLen := len(base)
	var digits []rune

	for nbr > 0 {
		digits = append(digits, rune(base[nbr%baseLen]))
		nbr /= baseLen
	}

	// Print in correct order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

// isValidBase checks if the base string is valid
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' || seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

// Helper function to print strings using z01.PrintRune
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
