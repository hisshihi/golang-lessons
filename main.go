package main

import "fmt"

type transformFn func(int) int

func main() {
	nubmers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	moreNumbers := []int{5, 34, 4, 22}
	doubleNumbers := transformNumbers(&nubmers, double)
	tripleNumbers := transformNumbers(&nubmers, triple)
	fmt.Printf("Double numbers: %v\n", doubleNumbers)
	fmt.Printf("Triple numbers: %v", tripleNumbers)

	transformerFu1 := getTransformFu(&nubmers)
	transformerFu2 := getTransformFu(&moreNumbers)

	transformedNumbers := transformNumbers(&nubmers, transformerFu1)
	moreTransformedNumbers := transformNumbers(&nubmers, transformerFu2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
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

func getTransformFu(numbers *[]int) transformFn {
	if (*numbers)[0] % 2 == 0 {
		return double
	} else {
		return triple
	}
}