package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, char := range word {
			z01.PrintRune(char) // print each character
		}
		z01.PrintRune('\n') // print newline after each word
	}
}
