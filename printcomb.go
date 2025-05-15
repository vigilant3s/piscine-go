package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ { // Iterate through the first digit from '0' to '9'.
		for j := i + 1; j <= '9'; j++ { // Iterate through the second digit from i+1 to '9'.
			for k := j + 1; k <= '9'; k++ { // Iterate through the third digit from j+1 to '9'.
				z01.PrintRune(i)                      // Print the first digit.
				z01.PrintRune(j)                      //	Print the second digit.
				z01.PrintRune(k)                      // Print the third digit.
				if i != '7' || j != '8' || k != '9' { // If the current combination is not the last one,
					z01.PrintRune(',') // Print a comma.
					z01.PrintRune(' ') // Print a space after the comma.
				}
			}
		}
	}
	z01.PrintRune('\n') // Print a newline at the end.
}
