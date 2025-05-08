package student

func UltimateDivMod(a *int, b *int) {
	if b != nil && *b != 0 { // Ensure b is not nil and not zero
		*a = *a / *b
		*b = *a % *b
	} else {
		// Handle division by zero and set a and b to zero
		*a = 0
		*b = 0
	}
}
