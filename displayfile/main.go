package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:] // Ignore program name

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(content))
}
