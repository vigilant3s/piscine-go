package piscine

func IsPrintable(s string) bool { // Check if all characters in the string are printable
	for _, r := range s { // Iterate over each rune in the string
		if r < 32 || r > 126 { // Check if the rune is outside the printable ASCII range
			return false // If any rune is not printable, return false
		}
	}
	return true // If all runes are printable, return true
}

// The function IsPrintable takes a string as input and checks if all characters in the string are printable ASCII characters.
// It returns true if all characters are printable and false otherwise.
