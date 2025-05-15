package piscine

func Swap(a *int, b *int) { // Swap swaps the values of two integers pointed to by a and b.
	*a, *b = *b, *a // Swap the values of the integers pointed to by a and b.
}
