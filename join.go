package piscine

func Join(strs []string, sep string) string { // Join concatenates a slice of strings with a separator
	result := ""               // Initialize an empty string to hold the result
	for i, str := range strs { // Iterate over the slice of strings
		if i > 0 { // If this is not the first element, add the separator
			result += sep // Add separator before all but the first element
		} // Add the current string to the result
		result += str // Add the string
	} // Return the concatenated result
	return result // Return the final concatenated string
}
