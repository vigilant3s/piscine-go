package piscine

func Unmatch(a []int) int {
	result := 0
	for _, num := range a {
		result ^= num
	}
	if result == 0 {
		return -1
	}
	return result
}
