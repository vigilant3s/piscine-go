package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) { // PrintWordsTables prints each word in the slice a on a new line.
	for _, word := range a { // iterate over each word in the slice
		for _, char := range word { // iterate over each character in the word
			z01.PrintRune(char) // print each character
		}
		z01.PrintRune('\n') // print newline after each word
	}
}
