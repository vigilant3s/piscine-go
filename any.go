package piscine

func Any(f func(string) bool, a []string) bool { // Any checks if any element in the slice satisfies the predicate function
	for i := 0; i < len(a); i++ { // Iterate through each element in the slice
		if f(a[i]) { // Check if the predicate function returns true for the current element
			return true // Return immediately if any match is found
		}
	}
	return false // None matched
}
