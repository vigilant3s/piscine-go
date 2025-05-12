package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}

// The function uses recursion to calculate the Fibonacci number at the given index.
// It checks the base cases for index being negative, zero, or one.
// If index is greater than one, it recursively calculates the Fibonacci number by summing the results of Fibonacci(index-1) and Fibonacci(index-2).
// The function returns the final result of the Fibonacci number at the specified index.
