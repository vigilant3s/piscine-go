package piscine

func IsUpper(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false // Found a non-uppercase character
		}
	}
	return true // All characters were uppercase
}
