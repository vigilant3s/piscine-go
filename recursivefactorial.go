package piscine

func RecursiveFactorial(arg int) int {
	if arg < 0 || arg > 20 {
		return 0
	}
	if arg == 0 || arg == 1 {
		return 1
	}
	return arg * RecursiveFactorial(arg-1)
}
