package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	digits := make([]int, 10)
	for n > 0 {
		d := n % 10
		digits[d]++
		n /= 10
	}
	for i := range digits {
		if digits[i] > 0 {
			for j := 0; j < digits[i]; j++ {
				z01.PrintRune(rune('0' + i))
			}
		}
	}
}
