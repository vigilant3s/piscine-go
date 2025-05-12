package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	return nb * RecursivePower(nb, power-1)
}

// The function uses recursion to calculate the power of nb.
// It checks the base cases for power being negative, zero, or one.
// If power is greater than one, it multiplies nb by the result of RecursivePower with power decremented by one.
// The function returns the final result of nb raised to the power of power.
// The time complexity of this function is O(n), where n is the value of power.
// The space complexity is O(n) due to the recursive call stack.
// The function handles negative powers by returning 0, as negative powers are not defined in this context.
