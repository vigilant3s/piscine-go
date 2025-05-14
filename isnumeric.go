package piscine

func IsNumeric(s string) bool { // Check if the string contains only digits
	for _, r := range s { // Iterate over each rune in the string
		if r < '0' || r > '9' { // Check if the rune is not a digit
			return false // If a non-digit is found, return false
		}
	}
	return true // If all runes are digits, return true
}
