package piscine

// AppendRange returns a slice of ints from min (inclusive) to max (exclusive)
func AppendRange(min, max int) []int {
	// If min is greater than or equal to max, return nil
	if min >= max {
		return nil
	}

	var result []int // Start with a nil slice

	// Loop from min to max (excluding max)
	for i := min; i < max; i++ {
		result = append(result, i) // Add each number to the slice
	}

	return result // Return the completed slice
}
