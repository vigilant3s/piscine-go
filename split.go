package piscine

func Split(s, sep string) []string { // Split splits a string s into substrings separated by sep.
	if sep == "" { // If sep is empty, return the original string in a slice.
		return []string{s} // Return the original string in a slice.
	}

	var result []string // Initialize an empty slice to hold the substrings.
	sepLen := len(sep)  // Get the length of the separator.
	lastIndex := 0      // Initialize the last index to 0.

	for i := 0; i <= len(s)-sepLen; { // Iterate through the string s.
		if s[i:i+sepLen] == sep { // If the substring from i to i+sepLen matches sep,
			result = append(result, s[lastIndex:i]) // add the substring from lastIndex to i to the result slice.
			i += sepLen                             //	 Move the index i forward by the length of sep.
			lastIndex = i                           // Update lastIndex to the current index i.
		} else { //	 If the substring does not match sep,
			i++ //	 increment i by 1.
		}
	}

	// Add remaining part after last separator
	result = append(result, s[lastIndex:]) // Add the remaining part of the string after the last separator to the result slice.

	return result // Return the slice of substrings
}
