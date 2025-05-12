package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb [10]int
	printed := false
	generateComb(comb, 0, 0, n, &printed)

	z01.PrintRune('\n')
}

func generateComb(comb [10]int, index, start, n int, printed *bool) {
	for i := start; i <= 10-(n-index); i++ {
		if i > 9 {
			break
		}
		comb[index] = i
		if index == n-1 {
			if *printed {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			for j := 0; j < n; j++ {
				z01.PrintRune(rune(comb[j] + '0'))
			}
			*printed = true
		} else {
			generateComb(comb, index+1, i+1, n, printed)
		}
	}
}
