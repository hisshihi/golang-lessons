package main

import "fmt"

type Product struct {
	price float64
	name  string
}

func printStats(list [4]Product) {
	var cost, totalItems float64
	for _, product := range list {
		cost += product.price
		if product.name != "" {
			totalItems++
		}
	}
	average := cost / totalItems
	fmt.Printf("Total cost: %.2f\n", cost)
	fmt.Printf("Average price: %.2f\n", average)
	fmt.Printf("Total items: %.0f\n", totalItems)
	fmt.Println("Last item:", list[len(list)-1].name)
}

func main() {
	shoppingList := [4]Product{
		{price: 1.99, name: "Milk"},
		{price: 0.99, name: "Eggs"},
		{price: 2.49, name: "Bread"},
	}

	shoppingList[3] = Product{price: 3.49, name: "Butter"}
	printStats(shoppingList)
}
