package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}

	// Start iterating from 1 upwards
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			return i // Return i if the square matches
		}
	}

	// Return 0 if no exact square root was found
	return 0
}

// The function Sqrt takes an integer nb as input.
// It calculates the integer square root of nb using a recursive approach.
// If nb is less than 0, it returns 0.
// If nb is 0 or 1, it returns nb.
// If nb is greater than 1, it recursively calls itself with nb-1 until it finds the integer square root.
