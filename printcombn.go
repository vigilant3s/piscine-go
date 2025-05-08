package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb [10]int
	generateComb(comb, 0, 0, n)
}

func generateComb(comb [10]int, index, start, n int) {
	if index == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(comb[i] + '0'))
		}
		if comb[0] != 10-n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for digit := start; digit <= 9; digit++ {
		// Avoid placing too large digits that would prevent full length n
		if 10-digit < n-index {
			break
		}
		comb[index] = digit
		generateComb(comb, index+1, digit+1, n)
	}
}
