package piscine

func UltimateDivMod(a *int, b *int) {
	if b != nil && *b != 0 { // Ensure b is not nil and not zero
		originalA := *a             // Store the original value of a
		*a = originalA / *b         // Perform integer division
		remainder := originalA % *b // Calculate the remainder
		*b = remainder              // Set b to the remainder
	} else {
		// Handle division by zero and set a and b to zero
		*a = 0 // Set a to zero
		*b = 0 // Set b to zero
	}
}
