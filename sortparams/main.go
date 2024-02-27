package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	array := os.Args[1:] // ["Husain", "Ali", "Fatima"]
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	for i := 0; i < len(array); i++ {
		for _, ch := range array[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
