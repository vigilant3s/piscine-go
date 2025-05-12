package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}

// The function IterativePower takes two integers nb and power as input.
// It calculates the value of nb raised to the power of power using an iterative approach.
// If power is negative, it returns 0.
