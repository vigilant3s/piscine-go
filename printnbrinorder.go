package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Special case: print 0 and return
	if n == 0 { // check if n is 0
		z01.PrintRune('0') // print 0 as a rune
		return             // return to avoid further processing
	}

	// Step 1: Extract digits
	var digits []int // slice to hold digits
	for n > 0 {      // while n is greater than 0
		d := n % 10                // get the last digit
		digits = append(digits, d) // append it to the slice
		n = n / 10                 // remove the last digit
	}

	// Step 2: Sort digits manually (simple bubble sort)
	for i := 0; i < len(digits); i++ { // iterate through digits
		for j := i + 1; j < len(digits); j++ { // compare digits
			if digits[j] < digits[i] { // if the next digit is smaller
				digits[i], digits[j] = digits[j], digits[i] // swap them
			}
		}
	}

	// Step 3: Print digits in ascending order
	for _, d := range digits { // iterate through sorted digits
		z01.PrintRune(rune(d + '0')) // convert int to rune character
	}
}
