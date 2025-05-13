package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r // First rune found, return it!
	}
	return 0 // If the string is empty
}

// The function returns 0, which is the zero value for rune type in Go.
