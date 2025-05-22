package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	if len(str) < 3 {
		z01.PrintRune('\n')
		return "\n"
	}

	result := []rune{}
	for i, r := range str {
		// Every third character means index 2, 5, 8, ... (0-based)
		if (i+1)%3 == 0 {
			result = append(result, r)
		}
	}

	// Print result runes
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	return string(result) + "\n"
}
