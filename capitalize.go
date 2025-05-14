package piscine

func Capitalize(s string) string { // Capitalize function
	runes := []rune(s) // Convert string to rune slice for proper handling of Unicode characters
	newWord := true    // Flag to indicate the start of a new word

	for i, r := range runes { // Iterate over each rune in the string
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') { // Check if the rune is a letter or digit
			if newWord { // If it's the start of a new word
				// If it's a lowercase letter, convert to uppercase
				if r >= 'a' && r <= 'z' { // Convert to uppercase
					runes[i] = r - 32 // ASCII conversion
				}
			} else { // If it's not the start of a new word
				// If it's an uppercase letter, convert to lowercase
				if r >= 'A' && r <= 'Z' { // Convert to lowercase
					runes[i] = r + 32 // ASCII conversion
				}
			}
			newWord = false //  Next character is not the start of a new word
		} else {
			newWord = true // next char starts a new word
		}
	}
	return string(runes) // Convert the rune slice back to a string and return it
}
