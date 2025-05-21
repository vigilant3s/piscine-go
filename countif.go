package piscine

func CountIf(f func(string) bool, tab []string) int { // CountIf counts the number of elements in the slice that satisfy the predicate function
	count := 0                      // Start with 0 valid elements
	for i := 0; i < len(tab); i++ { // Iterate through each string in the slice
		if f(tab[i]) { // Apply the function to each string
			count++ // If true, increment the count
		}
	}
	return count // Return the total number of true results
}
