package piscine

func Abort(a, b, c, d, e int) int {
	// Step 1: Put all values into an array
	nums := [5]int{a, b, c, d, e}

	// Step 2: Manually sort the array (Selection Sort)
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// Swap nums[i] and nums[minIndex]
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	// Step 3: Return the middle value (index 2)
	return nums[2]
}
