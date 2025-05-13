package piscine

func Compare(a, b string) int {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	return 0
}

// The Compare function compares two strings lexicographically.
// It returns -1 if the first string is less than the second string,
// 1 if the first string is greater than the second string, and 0 if they are equal.
// The function first determines the minimum length of the two strings.
// It then compares the characters of both strings one by one.
// If a character in the first string is less than the corresponding character in the second string,
// it returns -1. If it is greater, it returns 1.
// If all characters are equal, it checks the lengths of the strings.
