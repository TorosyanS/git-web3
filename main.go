package main

import (
	"fmt"
)

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	fmt.Println("Привет!")
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}
