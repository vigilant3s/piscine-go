package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb [10]int
	printed := false
	generateComb(comb, 0, 0, n, &printed)

	// ✅ Add final newline
	z01.PrintRune('\n')
}

func generateComb(comb [10]int, index, start, n int, printed *bool) {
	if index == n {
		if *printed {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(comb[i] + '0'))
		}
		*printed = true
		return
	}

	for i := start; i <= 9-(n-index); i++ {
		comb[index] = i
		generateComb(comb, index+1, i+1, n, printed)
	}
}
