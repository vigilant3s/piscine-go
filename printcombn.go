package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) { // PrintCombN prints all combinations of n digits in ascending order.
	if n <= 0 || n >= 10 { // Check if n is out of bounds (0 < n < 10)
		return // Return if n is invalid
	}

	var comb [10]int                      // Declare an array to hold the current combination
	printed := false                      // Flag to check if any combination has been printed
	generateComb(comb, 0, 0, n, &printed) // Call the recursive function to generate combinations

	z01.PrintRune('\n') // Print newline after all combinations are printed
}

func generateComb(comb [10]int, index, start, n int, printed *bool) { // generateComb generates combinations of n digits recursively.
	for i := start; i <= 10-(n-index); i++ { // Iterate from start to 10 - (n - index)
		if i > 9 { // If i exceeds 9, break the loop
			break // Exit the loop
		}
		comb[index] = i   // Assign the current index to i
		if index == n-1 { // If the current index is n-1, we have a complete combination
			if *printed { // If this is not the first combination printed,
				z01.PrintRune(',') // print a comma
				z01.PrintRune(' ') // and a space
			}
			for j := 0; j < n; j++ { // Iterate through the combination
				z01.PrintRune(rune(comb[j] + '0')) // Print each digit as a rune
			}
			*printed = true // Set the printed flag to true
		} else { // If the combination is not complete,
			generateComb(comb, index+1, i+1, n, printed) // recursively call generateComb with the next index and i+1
		}
	}
}
