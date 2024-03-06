package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			z01.PrintRune('0' + rune(i/10))
			z01.PrintRune('0' + rune(i%10))
			z01.PrintRune(' ')

			z01.PrintRune('0' + rune(j/10))
			z01.PrintRune('0' + rune(j%10))

			if !(i == 1 && j == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
