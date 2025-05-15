package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// Calculate total size: sum of word lengths + number of newlines (len - 1)
	totalLen := 0
	for _, word := range args { // Calculate the length of each word
		totalLen += len(word) // Add the length of the word to the total length
	}
	totalLen += len(args) - 1 // for "\n"

	// Create a byte slice to build the string
	result := make([]byte, 0, totalLen)

	for i, word := range args { // Iterate over each word
		for j := 0; j < len(word); j++ { // Iterate over each character in the word
			result = append(result, word[j]) // Append the character to the result
		}
		if i < len(args)-1 { // If not the last word, append a newline
			result = append(result, '\n') // Append a newline character
		}
	}

	return string(result) // Convert the byte slice to a string and return it
}
