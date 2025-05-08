package student

func DivMod(a int, b int, div *int, mod *int) {
	if b != 0 {
		*div = a / b // Perform division if b is not 0
		*mod = a % b // Perform modulo if b is not 0
	} else {
		// Handle the case where division by zero occurs
		*div = 0 // Set the result of div to 0
		*mod = 0 // Set the result of mod to 0
	}
}
