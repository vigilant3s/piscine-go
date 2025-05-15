package piscine

import "strings"

// ConcatParams joins a slice of strings with newline characters
func ConcatParams(args []string) string {
	// Join all strings in the slice with "\n" between them
	return strings.Join(args, "\n")
}
