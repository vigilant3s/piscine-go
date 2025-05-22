package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	// Step 1: Put all inputs into a slice
	nums := []int{a, b, c, d, e}

	// Step 2: Sort the slice in ascending order
	sort.Ints(nums)

	// Step 3: Return the middle element (index 2)
	return nums[2]
}
