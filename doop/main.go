package main

import (
	"os"
)

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}

func printInt(n int64) {
	if n == 0 {
		printStr("0\n")
		return
	}
	if n < 0 {
		printStr("-")
		n = -n
	}
	var digits [20]byte
	i := len(digits)
	for n > 0 {
		i--
		digits[i] = byte(n%10) + '0'
		n /= 10
	}
	os.Stdout.Write(digits[i:])
	os.Stdout.Write([]byte{'\n'})
}

// parseInt64 parses string to int64, returns value and success bool.
func parseInt64(s string) (int64, bool) {
	var n int64
	if len(s) == 0 {
		return 0, false
	}
	i := 0
	neg := false
	if s[0] == '-' {
		neg = true
		i++
		if i == len(s) {
			return 0, false
		}
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := int64(c - '0')
		// Check overflow before multiplying by 10
		if n > (1<<63-1)/10 {
			return 0, false
		}
		n = n*10 + d
	}
	if neg {
		n = -n
	}
	return n, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := parseInt64(os.Args[1])
	op := os.Args[2]
	b, ok2 := parseInt64(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		// Overflow check for addition
		if (b > 0 && a > (1<<63-1)-b) || (b < 0 && a < (-1<<63)-b) {
			return
		}
		printInt(a + b)
	case "-":
		// Overflow check for subtraction
		if (b < 0 && a > (1<<63-1)+b) || (b > 0 && a < (-1<<63)+b) {
			return
		}
		printInt(a - b)
	case "*":
		// multiplication overflow check
		if a == 0 || b == 0 {
			printInt(0)
			return
		}
		absA := a
		if absA < 0 {
			absA = -absA
		}
		absB := b
		if absB < 0 {
			absB = -absB
		}
		if absA > (1<<63-1)/absB {
			return
		}
		printInt(a * b)
	case "/":
		if b == 0 {
			printStr("No division by 0\n")
			return
		}
		printInt(a / b)
	case "%":
		if b == 0 {
			printStr("No modulo by 0\n")
			return
		}
		printInt(a % b)
	default:
		return
	}
}
