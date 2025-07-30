package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumup(10, 2, 3, 4, 5, 40, 234, 11))
}

func sumup(startingValue int, numbers ...int) int {
	sumNumber := startingValue

	for _, number := range numbers {
		sumNumber += number
	}
	return sumNumber
}
