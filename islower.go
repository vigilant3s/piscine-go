package piscine

func IsLower(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false // Not a lowercase letter
		}
	}
	return true // All letters are lowercase
}

// This function checks if all characters in the string are lowercase letters.
// It returns true if all characters are lowercase letters, and false otherwise.
// The function iterates over each character in the string and checks if it falls.
