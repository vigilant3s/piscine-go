package piscine

func ToUpper(s string) string { // Convert a string to uppercase
	result := ""          // Initialize an empty string to store the result
	for _, r := range s { // Iterate over each rune in the string
		if r >= 'a' && r <= 'z' { // Check if the rune is a lowercase letter
			r = r - 32 // Convert to uppercase
		} // Else, it remains unchanged
		result += string(r) // Append the rune to the result string
	} // End of the loop
	return result // Return the final uppercase string
} // End of the function
