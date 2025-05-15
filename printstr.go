package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) { // PrintStr prints the string s character by character.
	for _, char := range s { // Iterate over each character in the string
		z01.PrintRune(char) // Print each character
	}
}
