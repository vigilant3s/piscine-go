package student

func UltimateDivMod(a *int, b *int) {
	if *b != 0 {
		*a = *a / *b
		*b = *a % *b
	} else {
		*a = 0
		*b = 0
	}
}
