package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 { // Edge case for min int on 64-bit (if needed)
		z01.PrintRune('-')           // Print negative sign
		PrintNbr(922337203685477580) // Print the absolute value
		z01.PrintRune('8')           // Print the last digit
		return                       // Return to avoid further recursion
	}

	if n < 0 { // If n is negative, print the negative sign and convert n to positive
		z01.PrintRune('-') // Print negative sign
		n = -n             // Convert n to positive
	}

	if n >= 10 { // If n is greater than or equal to 10, recursively call PrintNbr with n divided by 10
		PrintNbr(n / 10) // Recursive call with n divided by 10
	}
	z01.PrintRune(rune('0' + n%10)) // Print the last digit of n
}
