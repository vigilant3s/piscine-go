package piscine

func LastRune(s string) rune {
	var last rune
	for _, r := range s {
		last = r
	}
	return last
}

// // The function returns the first rune of the string.
// // Compare this snippet from lastrune_test.go:
