package piscine

func Unmatch(a []int) int {
	freq := make(map[int]int) // Step 1: Count frequency of each number

	for _, num := range a {
		freq[num]++
	}

	// Step 2: Loop in order to find the first odd-count number
	for _, num := range a {
		if freq[num]%2 != 0 {
			return num
		}
	}

	return -1 // All numbers had even count
}
