package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	var insertStr string
	var baseStr string
	shouldSort := false

	// Parse flags and arguments manually
	for i := 0; i < len(args); i++ {
		arg := args[i]

		// Handle --insert or -i
		if (len(arg) > 8 && arg[:8] == "--insert") || (len(arg) > 2 && arg[:2] == "-i") {
			// Check if the argument has an '=' to separate the flag and value
			if len(arg) > 8 && arg[8] == '=' {
				insertStr = arg[9:] // After '='
			} else if len(arg) > 2 && arg[2] == '=' {
				insertStr = arg[3:] // After '='
			}
		} else if arg == "--order" || arg == "-o" {
			shouldSort = true
		} else if arg[0] != '-' {
			baseStr = arg // The first argument that is not a flag
		}
	}

	// Combine base string and insert string
	finalStr := baseStr + insertStr

	// If --order flag is used, sort the string manually
	if shouldSort {
		finalStr = manualSort(finalStr)
	}

	// Print the final string character by character
	for _, r := range finalStr {
		z01.PrintRune(r)
	}
}

func printHelp() {
	lines := []string{
		"--insert",
		"  -i",
		"\t This flag inserts the string into the string passed as argument.",
		"--order",
		"  -o",
		"\t This flag will behave like a boolean, if it is called it will order the argument.",
	}
	for _, line := range lines {
		for _, r := range line {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

// Manual sorting function (ASCII order)
func manualSort(s string) string {
	// Convert string to a slice of runes (characters)
	runes := []rune(s)
	// Simple bubble sort algorithm
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-1-i; j++ {
			if runes[j] > runes[j+1] {
				// Swap the characters
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	// Convert back to string
	return string(runes)
}
