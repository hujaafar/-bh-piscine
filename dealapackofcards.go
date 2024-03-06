package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	cardsPerPlayer := len(deck) / 4
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < cardsPerPlayer; j++ {
			if j != 0 {
				fmt.Print(", ")
			}
			fmt.Print(deck[i*cardsPerPlayer+j])
		}
		fmt.Print("\n")
	}
}
