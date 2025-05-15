package piscine

import "github.com/01-edu/z01"

func PrintComb2() { // PrintComb2 prints all combinations of two different digits in ascending order.
	for i := 0; i <= 98; i++ { // Iterate through the first digit from 0 to 98.
		for j := i + 1; j <= 99; j++ { // Iterate through the second digit from i+1 to 99.
			z01.PrintRune(rune('0' + i/10)) // Print the first digit.
			z01.PrintRune(rune('0' + i%10)) // Print the second digit.
			z01.PrintRune(' ')              // Print a space between the two digits.
			z01.PrintRune(rune('0' + j/10)) // Print the third digit.
			z01.PrintRune(rune('0' + j%10)) // Print the fourth digit.

			if i != 98 || j != 99 { // If the current combination is not the last one,
				z01.PrintRune(',') // Print a comma.
				z01.PrintRune(' ') // Print a space after the comma.
			}
		}
	}
	z01.PrintRune('\n') // Print a newline at the end.
}
