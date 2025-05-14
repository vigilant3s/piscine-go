package piscine

func Index(s string, toFind string) int {
	// Convert both strings to rune slices to handle all characters properly
	runes := []rune(s)
	find := []rune(toFind)

	// Get lengths of both rune slices
	lenRunes := len(runes)
	lenFind := len(find)

	// If toFind is empty, or longer than the main string, return -1
	if lenFind == 0 || lenFind > lenRunes {
		return -1
	}

	// Loop through each index where toFind could possibly start
	for i := 0; i <= lenRunes-lenFind; i++ {
		match := true

		// Check if characters starting from i match toFind
		for j := 0; j < lenFind; j++ {
			if runes[i+j] != find[j] {
				match = false
				break
			}
		}

		// If we found a match, return the index
		if match {
			return i
		}
	}

	// If nothing matched
	return -1
}
