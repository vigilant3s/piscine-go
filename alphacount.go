package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}

// AlphaCount counts the number of alphabetic characters in a string.
// It iterates through each character in the string and checks if it is an alphabetic character.
// If it is, it increments the count.
// Finally, it returns the total count of alphabetic characters.
// The function uses a for loop to iterate through the string and a conditional statement to check if each character is alphabetic.
// The function returns the total count of alphabetic characters in the string.
