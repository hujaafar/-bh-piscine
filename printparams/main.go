package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args
	arr2 := []rune(arr[0])

	for i := 1; i < len(arr2); i++ {
		z01.PrintRune(arr2[i])
	}
	z01.PrintRune('\n')
}
