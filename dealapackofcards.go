package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	player := 1
	for j := 0; j < 12; {
		fmt.Printf("Player %d: ", player)
		for i := 1; i < 4; i++ {
			fmt.Printf("%d", deck[j])
			j++
			if i != 3 {
				fmt.Printf(", ")
			}
		}
		player++
		fmt.Printf("\n")
	}
}
