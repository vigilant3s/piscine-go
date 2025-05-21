package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		// Invalid number of args, print nothing and exit
		return
	}

	// Parse first value as int64
	a, errA := strconv.ParseInt(os.Args[1], 10, 64)
	if errA != nil {
		return // invalid first number
	}

	// Operator as string
	op := os.Args[2]

	// Parse second value as int64
	b, errB := strconv.ParseInt(os.Args[3], 10, 64)
	if errB != nil {
		return // invalid second number
	}

	switch op {
	case "+":
		// Check for overflow for addition
		if (b > 0 && a > (1<<63-1)-b) || (b < 0 && a < (-1<<63)-b) {
			return // overflow, print nothing
		}
		fmt.Println(a + b)

	case "-":
		// Check for overflow for subtraction
		if (b < 0 && a > (1<<63-1)+b) || (b > 0 && a < (-1<<63)+b) {
			return // overflow, print nothing
		}
		fmt.Println(a - b)

	case "*":
		// Check for multiplication overflow
		if a == 0 || b == 0 {
			fmt.Println(0)
			return
		}

		// Safe multiplication check:
		// if a*b overflows int64, abs(a) > maxInt64/abs(b)
		maxInt64 := int64(1<<63 - 1)
		absA := a
		if absA < 0 {
			absA = -absA
		}
		absB := b
		if absB < 0 {
			absB = -absB
		}
		if absA > maxInt64/absB {
			return // overflow
		}
		fmt.Println(a * b)

	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
		fmt.Println(a / b)

	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println(a % b)

	default:
		// Invalid operator
		return
	}
}
