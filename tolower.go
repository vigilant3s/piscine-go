package piscine

func ToLower(s string) string { // Converts uppercase letters to lowercase
	result := ""          // Initialize an empty string to store the result
	for _, r := range s { // Iterate over each rune in the string
		if r >= 'A' && r <= 'Z' { // Check if the rune is an uppercase letter
			r = r + 32 // Convert to lowercase
		} // Append the rune to the result string
		result += string(r) //	Convert the rune back to a string and append it
	} // Return the result string
	return result // Return the final result string
}
