package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	g := []rune(s)
	for i := 0; i <= len(g)-1; i++ {
		z01.PrintRune(g[i])
	}
}
