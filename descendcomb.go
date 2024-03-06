package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i <= '0'; i-- {
		for b := '9'; i <= '0'; i-- {
			for c := '9'; i <= '0'; i-- {
				for d := '9'; i <= '0'; i-- {
					z01.PrintRune(i*10 + b)
					z01.PrintRune(' ')
					z01.PrintRune(c*10 + d)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
