package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args

	for i := len(arr) - 1; i >= 0; i-- {
		if i > 0 {
			for _, chr := range arr[i] {
				z01.PrintRune(chr)
			}
			z01.PrintRune('\n')
		}
	}
}
