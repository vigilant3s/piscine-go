package piscine

func StrRev(s string) string { // StrRev reverses the string s and returns the reversed string.
	runes := []rune(s)                                    // Convert the string s to a slice of runes.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // Iterate through the slice of runes, swapping elements from both ends.
		runes[i], runes[j] = runes[j], runes[i] // Swap the elements at indices i and j.
	}
	return string(runes) // Convert the slice of runes back to a string and return it.
}
