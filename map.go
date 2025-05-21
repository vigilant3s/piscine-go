package piscine

func Map(f func(int) bool, a []int) []bool { // Map applies function f to each element of slice a
	result := make([]bool, len(a)) // Create a bool slice of same length as 'a'
	for i := 0; i < len(a); i++ {  // Iterate over each element in 'a'
		result[i] = f(a[i]) // Apply f to each element, store result
	}
	return result // Return the resulting slice
} // End of Map function
