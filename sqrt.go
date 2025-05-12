package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return nb
	}
	return Sqrt(nb - 1)
}

// The function Sqrt takes an integer nb as input.
// It calculates the integer square root of nb using a recursive approach.
// If nb is less than 0, it returns 0.
// If nb is 0 or 1, it returns nb.
// If nb is greater than 1, it recursively calls itself with nb-1 until it finds the integer square root.
