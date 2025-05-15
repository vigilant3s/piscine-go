package piscine

func SortIntegerTable(table []int) { // SortIntegerTable sorts a slice of integers in ascending order using the bubble sort algorithm.
	n := len(table) // Get the length of the slice.
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ { // Iterate through the slice, comparing adjacent elements.
			if table[j] > table[j+1] { // If the current element is greater than the next element,
				table[j], table[j+1] = table[j+1], table[j] // swap them.
			}
		}
	}
}
