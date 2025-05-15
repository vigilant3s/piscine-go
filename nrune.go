package piscine

func NRune(s string, n int) rune { // NRune returns the nth rune of the string s.
	if n <= 0 { // If n is less than or equal to 0, return 0.
		return 0 // Return 0.
	}

	count := 1            // Initialize a counter to 1.
	for _, r := range s { // Iterate over the string s, counting the runes.
		if count == n { // If the counter matches n,
			return r //	 return the current rune.
		}
		count++ // Increment the counter.
	}
	return 0 //	 If the nth rune is not found, return 0.
}

// NRune returns the nth rune of the string s. If n is less than or equal to 0, it returns 0.
// If n is greater than the length of s, it returns 0.
// The function iterates over the string s, counting the runes until it reaches the nth one.
// It uses a for loop to iterate over the string and a counter to keep track of the current rune index.
// If the nth rune is found, it returns that rune.
// If the end of the string is reached without finding the nth rune, it returns 0.
// The function handles multi-byte runes correctly by using the range keyword in the for loop.
// It is important to note that the function does not modify the original string s.
// The function is useful for extracting a specific rune from a string, which can be helpful in various string manipulation tasks.
