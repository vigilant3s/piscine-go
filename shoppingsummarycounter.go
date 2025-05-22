package piscine

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	// Create a map to store item counts
	counts := make(map[string]int)

	// Split the string into items using space as separator
	items := strings.Split(str, " ")

	// Count each item
	for _, item := range items {
		counts[item]++
	}

	return counts
}
