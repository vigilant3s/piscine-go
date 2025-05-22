package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)
	word := ""

	for _, char := range str {
		if char == ' ' {
			if word != "" {
				counts[word]++
				word = ""
			}
			// If word == "", just skip (multiple spaces)
		} else {
			word += string(char)
		}
	}

	// Add the last word if there's any
	if word != "" {
		counts[word]++
	}

	return counts
}
