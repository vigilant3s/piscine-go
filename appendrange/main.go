package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(5, 10)) // Output: [5 6 7 8 9]
	fmt.Println(piscine.AppendRange(10, 5)) // Output: []
	fmt.Println(piscine.AppendRange(0, 1))  // Output: [0]
	fmt.Println(piscine.AppendRange(-3, 2)) // Output: [-3 -2 -1 0 1]
	fmt.Println(piscine.AppendRange(0, 0))  // Output: []
}
