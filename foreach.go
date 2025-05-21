package piscine

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i]) // Apply the function to each element
	}
}
