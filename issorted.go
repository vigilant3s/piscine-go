package piscine

func IsSorted(f func(a, b int) int, a []int) bool { // IsSorted checks if the slice is sorted according to the provided comparison function
	if len(a) < 2 { // If the slice has less than 2 elements, it's considered sorted
		return true // Empty or one-element slice is always sorted
	}

	// Determine initial direction (ascending or descending)
	ascending := false
	descending := false

	for i := 0; i < len(a)-1; i++ { // Iterate through the slice
		result := f(a[i], a[i+1]) // Compare adjacent elements using the provided function

		if result > 0 { // If the first element is greater than the second
			descending = true // Set descending to true
		} else if result < 0 { // If the first element is less than the second
			ascending = true // Set ascending to true
		}

		// If both ascending and descending detected, it's not sorted
		if ascending && descending { // Return false
			return false
		}
	}
	return true // If we finish the loop without finding both directions, it's sorted
}
