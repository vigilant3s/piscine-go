package piscine

func SplitWhiteSpaces(s string) []string { // SplitWhiteSpaces splits a string into substrings separated by whitespace characters.
	var result []string // Initialize an empty slice to hold the substrings
	length := len(s)    // Get the length of the input string
	start := -1         // Initialize start index to -1, indicating no word has been found yet

	for i := 0; i < length; i++ { // Iterate through each character in the string
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' { // Check if the character is not a whitespace
			if start == -1 { //	If we are not currently in a word
				start = i // Set the start index to the current position
			}
		} else { // If we encounter a whitespace character
			if start != -1 { // If we are currently in a word
				result = append(result, s[start:i]) // Append the word to the result slice
				start = -1                          // Reset start index to -1, indicating we are not in a word anymore
			}
		}
	}

	// Catch final word if the string didn't end with a separator
	if start != -1 { // If we are still in a word at the end of the string
		result = append(result, s[start:]) // Append the last word to the result slice
	}

	return result // Return the slice of substrings
}
