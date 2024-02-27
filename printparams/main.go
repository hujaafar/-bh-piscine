package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args

	for i := range arr {
		if i > 0 {
			for _, chr := range arr[i] {
				z01.PrintRune(chr)
			}
			z01.PrintRune('\n')
		}
	}
}
