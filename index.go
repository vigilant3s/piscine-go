package piscine

func Index(s string, toFind string) int {
	// Convert both strings to rune slices to handle Unicode characters correctly
	runes := []rune(s)     // Convert the main string to a slice of runes
	find := []rune(toFind) // Convert the string to find to a slice of runes

	// Get the lengths of both rune slices
	lenRunes := len(runes) // Length of the main string
	lenFind := len(find)   //	 Length of the string to find

	// ✅ Special case: if toFind is an empty string, return 0
	// By convention, an empty string is considered found at the start of any string
	if lenFind == 0 { // If the string to find is empty, return 0
		return 0 //	 Return 0
	}

	// If the string to find is longer than the main string, it can't be present
	if lenFind > lenRunes { // If the string to find is longer than the main string,
		return -1 // return -1
	}

	// Loop through each index in the main string where toFind could start
	for i := 0; i <= lenRunes-lenFind; i++ {
		match := true // Assume it's a match unless proven otherwise

		// Check each character of toFind against runes starting at index i
		for j := 0; j < lenFind; j++ { // Loop through each character of the string to find
			if runes[i+j] != find[j] { // If the character at index i+j in the main string does not match the character at index j in the string to find,
				match = false // Characters don't match — break out
				break         //
			}
		}

		// If we went through the whole toFind and it matched, return the start index
		if match { // If we found a match,
			return i // return the index where the match starts
		}
	}

	// If we finish the loop without finding a match, return -1
	return -1
}
