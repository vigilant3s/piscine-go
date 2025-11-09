// go_reloaded/main.go
package main

import (
	"fmt"
	"os"
	"piscine/go_reloaded/textfix"
)

func main() {
	// Expect exactly two args: input and output file names.
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inFile := os.Args[1]
	outFile := os.Args[2]

	// 1) Read input file
	inputBytes, err := os.ReadFile(inFile)
	if err != nil {
		fmt.Printf("read error: %v\n", err)
		return
	}

	// 2) Process the text through our pipeline
	output := textfix.ProcessText(string(inputBytes))

	// 3) Write output file
	if err := os.WriteFile(outFile, []byte(output), 0644); err != nil {
		fmt.Printf("write error: %v\n", err)
	}
}
