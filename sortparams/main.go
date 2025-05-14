package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Helper function to compare two strings in ASCII order
func compare(s1, s2 string) bool {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return true
		} else if s1[i] > s2[i] {
			return false
		}
	}
	// If all previous chars are equal, shorter string is smaller
	return len(s1) < len(s2)
}

// Helper function to print a string with z01
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	// Basic Bubble Sort
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-1-i; j++ {
			if !compare(args[j], args[j+1]) {
				// Swap args[j] and args[j+1]
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Print each sorted argument
	for _, s := range args {
		printStr(s)
	}
}
