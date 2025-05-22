package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < 3; j++ {
			card := deck[i*3+j]
			if card >= 10 {
				fmt.Printf("%d", card/10)
			}
			fmt.Printf("%d", card%10)
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
