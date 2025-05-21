package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if f(a[j], a[j+1]) > 0 {
				// Swap if out of order according to f
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
