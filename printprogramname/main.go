package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	maina := os.Args
	word := []rune(maina[0])

	for i := 0; i < len(word); i++ {
		if word[i] != '.' && word[i] != '/' {
			z01.PrintRune(word[i])
		}
	}
	z01.PrintRune('\n')
}
