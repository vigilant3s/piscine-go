package piscine

import "github.com/01-edu/z01"

// PrintNbrBase prints an integer in a custom base
func PrintNbrBase(nbr int, base string) {
	// Check if base is valid
	if !isValidBase(base) { // check for valid base
		printStr("NV") // Invalid base
		return
	}

	if nbr == 0 { // special case for 0
		z01.PrintRune(rune(base[0])) // print first character of base
		return                       // exit
	}

	// Handle negative
	if nbr < 0 { // if negative number
		z01.PrintRune('-') // print negative sign
		nbr = -nbr         // make it positive
	}

	// Convert to base
	baseLen := len(base) // length of base
	var digits []rune    // slice to hold digits

	for nbr > 0 { // convert number to base
		digits = append(digits, rune(base[nbr%baseLen])) // append the digit
		nbr /= baseLen                                   // divide by base
	}

	// Print in correct order
	for i := len(digits) - 1; i >= 0; i-- { // reverse order
		z01.PrintRune(digits[i]) // print each digit
	}
}

// isValidBase checks if the base string is valid
func isValidBase(base string) bool { // check if base is valid
	if len(base) < 2 { // base must have at least 2 characters
		return false // invalid base
	}

	seen := make(map[rune]bool) // map to track seen characters
	for _, ch := range base {   // iterate over each character
		if ch == '+' || ch == '-' || seen[ch] { // check for invalid characters
			return false // invalid base
		}
		seen[ch] = true // mark character as seen
	}
	return true // valid base
}

// Helper function to print strings using z01.PrintRune
func printStr(s string) { // print string
	for _, r := range s { // iterate over each character
		z01.PrintRune(r) // print character
	}
}
