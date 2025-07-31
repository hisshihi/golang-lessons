package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 10, 23, 84, 44, -4)
	anotherSum := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int {
	sumNumber := 0

	for _, number := range numbers {
		sumNumber += number
	}
		return sumNumber
}
