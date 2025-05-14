package piscine

func BasicJoin(elems []string) string { // BasicJoin concatenates all strings in the slice into a single string
	result := ""                // initialize an empty string to hold the result
	for _, str := range elems { // iterate over each string in the slice
		result += str // concatenate each string in the slice
	} // return the concatenated string
	return result // return the final concatenated string
}
