package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}
	
	transformedDouble := transformFu(numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformedDouble)

	transformedTriple := transformFu(numbers, func(number int) int {
		return number * 3
	})
	fmt.Println(transformedTriple)
}

func transformFu(numbers []int, transfrom func(int) int) []int {
	transformedNumbers := []int{}

	for _, number := range numbers {
		transformedNumbers = append(transformedNumbers, transfrom(number))
	}
	return transformedNumbers
}
