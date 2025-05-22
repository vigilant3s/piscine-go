package piscine

func Unmatch(a []int) int {
	freq := make(map[int]int) // Frequency map

	// Count the occurrences of each number
	for _, num := range a {
		freq[num]++
	}

	// Find the number that occurs an odd number of times
	for num, count := range freq {
		if count%2 != 0 {
			return num
		}
	}

	return -1
}
