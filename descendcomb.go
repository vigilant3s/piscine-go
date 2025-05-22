package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := true // Flag to track if it's the first combination or not

	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			// Print the first number
			z01.PrintRune(rune(i/10 + '0'))
			z01.PrintRune(rune(i%10 + '0'))
			z01.PrintRune(' ')

			// Print the second number
			z01.PrintRune(rune(j/10 + '0'))
			z01.PrintRune(rune(j%10 + '0'))

			// Add a comma and space if it's not the last combination
			if !(i == 1 && j == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	// Print newline after all combinations are printed, only once at the end
	z01.PrintRune('\n')
}
