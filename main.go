package main

import "fmt"

type transformFn func(int) int

func main() {
	nubmers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubleNumbers := transformNumbers(&nubmers, double)
	tripleNumbers := transformNumbers(&nubmers, triple)
	fmt.Printf("Double numbers: %v\n", doubleNumbers)
	fmt.Printf("Triple numbers: %v", tripleNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
