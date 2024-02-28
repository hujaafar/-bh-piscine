package main

import (
	"os"

	"github.com/01-edu/z01"
)

func check(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	slice := []rune{}
	result := ""
	count := 0
	IsF := true
	for _, c := range args {
		for _, j := range c {
			if check(j) {
				slice = append(slice, j)
				count++
			}
		}
		if IsF {
			result = c
			IsF = false
			continue
		}
		result = result + " " + c
	}
	cur := 0
	for _, c := range result {
		if check(c) {
			z01.PrintRune(slice[count-cur-1])
			cur++
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
