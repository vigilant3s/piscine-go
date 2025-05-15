package piscine

func RecursiveFactorial(arg int) int { // RecursiveFactorial calculates the factorial of a number using recursion.
	if arg < 0 || arg > 20 { // If arg is negative or greater than 20, return 0 as factorial is not defined for these cases.
		return 0 //	 Return 0 for invalid input.
	}
	if arg == 0 || arg == 1 { // If arg is 0 or 1, return 1 as the factorial of 0 and 1 is 1.
		return 1 // Return 1 for base case.
	}
	return arg * RecursiveFactorial(arg-1) // Recursively multiply arg by the result of RecursiveFactorial with arg decremented by 1.
}
