package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RecursivePower(4, 3))
}

// Output: 64
// The function RecursivePower takes two integers nb and power as input.
// It calculates the value of nb raised to the power of power using a recursive approach.
// If power is negative, it returns 0.
// If power is 0, it returns 1.
// If power is greater than 0, it recursively multiplies nb by the result of RecursivePower(nb, power-1).
// If power is 1, it returns nb.
