package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	found := false
	for _, v := range args {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			found = true
		}
	}
	if found {
		fmt.Println("Alert!!!")
	}
}
