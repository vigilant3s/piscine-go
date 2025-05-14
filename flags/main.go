package main

import (
	"os"
	"sort"
	"strings"

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

	for _, arg := range args {
		if strings.HasPrefix(arg, "--insert=") {
			insertStr = strings.TrimPrefix(arg, "--insert=")
		} else if strings.HasPrefix(arg, "-i=") {
			insertStr = strings.TrimPrefix(arg, "-i=")
		} else if arg == "--order" || arg == "-o" {
			shouldSort = true
		} else if !strings.HasPrefix(arg, "-") {
			baseStr = arg
		}
	}

	finalStr := baseStr + insertStr

	if shouldSort {
		runes := []rune(finalStr)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		finalStr = string(runes)
	}

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
