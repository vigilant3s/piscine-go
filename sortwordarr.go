package piscine

// SortWordArr sorts a slice of strings in ascending ASCII order using Bubble Sort.
func SortWordArr(a []string) { // Bubble Sort algorithm
	n := len(a)              // Get the length of the slice
	for i := 0; i < n; i++ { // Outer loop for n iterations
		for j := 0; j < n-1-i; j++ { // Inner loop for n-i-1 iterations
			if a[j] > a[j+1] { // Compare adjacent elements
				// Swap elements if out of order
				a[j], a[j+1] = a[j+1], a[j] // Swap the elements
			}
		}
	}
}
