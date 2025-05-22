package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)
	word := ""

	for _, char := range str {
		if char == ' ' {
			// count the word, even if empty
			counts[word]++
			word = ""
		} else {
			word += string(char)
		}
	}

	// add last word
	counts[word]++

	return counts
}
