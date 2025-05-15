package piscine

func StrLen(s string) int { // StrLen returns the length of the string s.
	count := 0    // Initialize a counter to 0.
	for range s { // Iterate through each character in the string s.
		count++ // Increment the counter for each character.
	}
	return count //	 Return the length of the string.
}
