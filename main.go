package main

import (
	"fmt"
)

func main() {
	c1 := makeCounter(10)
	c2 := makeCounter(15)

	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())
}

func makeCounter(number int) func() int {
	count := number
	return func() int {
		count++
		return count
	}
}
