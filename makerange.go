package piscine

// MakeRange returns a slice from min to max-1
func MakeRange(min, max int) []int {
	// If min is not less than max, return nil
	if min >= max {
		return nil
	}

	// Calculate the length of the slice
	length := max - min

	// Create a slice of ints with specified length
	result := make([]int, length)

	// Fill the slice with values from min to max-1
	for i := 0; i < length; i++ {
		result[i] = min + i
	}

	return result
}
