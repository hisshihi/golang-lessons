package main

import "fmt"

func main() {
	prices := []float64{10.99, 20.49}
	fmt.Printf("Prices: %v\n", prices)
	fmt.Printf("Second element in prices: %v\n", prices[1])

	prices = append(prices, 19.99)
	fmt.Printf("Updated prices: %v\n", prices)
}
