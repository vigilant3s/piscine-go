package student

func UltimateDivMod(a *int, b *int) {
	*a = *a / *b // Divide the value pointed to by 'a' by the value pointed to by 'b'
	*b = *a % *b // Store the remainder in the location pointed to by 'b'
}
