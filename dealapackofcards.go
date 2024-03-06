package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	players := map[int][]int{}
	for i, card := range deck {
		player := i % 4
		players[player] = append(players[player], card)
	}
	for i := 0; i < 4; i++ {
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')
		z01.PrintRune(rune('1' + i))
		z01.PrintRune(':')
		z01.PrintRune(' ')
		for j, card := range players[i] {
			printNumber(card)
			if j < len(players[i])-1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func printNumber(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printNumber(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}
