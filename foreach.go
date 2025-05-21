package piscine

func ForEach(f func(int), a []int) { // ForEach applies the function f to each element of the slice a
	for i := 0; i < len(a); i++ { // Iterate over the slice
		f(a[i]) // Apply the function to each element
	}
}
