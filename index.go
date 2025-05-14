package piscine

func Index(s string, toFind string) int {
	// Convert both strings to rune slices to handle Unicode characters correctly
	runes := []rune(s)
	find := []rune(toFind)

	// Get the lengths of both rune slices
	lenRunes := len(runes)
	lenFind := len(find)

	// ✅ Special case: if toFind is an empty string, return 0
	// By convention, an empty string is considered found at the start of any string
	if lenFind == 0 {
		return 0
	}

	// If the string to find is longer than the main string, it can't be present
	if lenFind > lenRunes {
		return -1
	}

	// Loop through each index in the main string where toFind could start
	for i := 0; i <= lenRunes-lenFind; i++ {
		match := true // Assume it's a match unless proven otherwise

		// Check each character of toFind against runes starting at index i
		for j := 0; j < lenFind; j++ {
			if runes[i+j] != find[j] {
				match = false // Characters don't match — break out
				break
			}
		}

		// If we went through the whole toFind and it matched, return the start index
		if match {
			return i
		}
	}

	// If we finish the loop without finding a match, return -1
	return -1
}
