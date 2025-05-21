package main

import (
	"github.com/01-edu/z01"
)

// Define the point struct with int fields
type point struct {
	x int
	y int
}

// Set x and y values via pointer
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Helper to print an int using z01
func printNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

// Print full message: x = 42, y = 21
func main() {
	points := &point{}
	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNbr(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNbr(points.y)
	z01.PrintRune('\n')
}
