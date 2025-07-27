package main

import "fmt"

func main() {
	type Product struct {
		Title string
		ID    int64
		Price float64
	}

	products := []Product{
		{Title: "Laptop", ID: 1, Price: 999.99},
		{Title: "Smartphone", ID: 2, Price: 499.99},
	}

	discountProducts := []Product{
		{Title: "Iphone 16e", ID: 3, Price: 699.99},
		{Title: "Loffree", ID: 4, Price: 299.99},
		{Title: "Logiteck mx master3s", ID: 5, Price: 199.99},
		{Title: "MacBook Pro", ID: 6, Price: 1299.99},
	}

	products = append(products, discountProducts...)
	for _, product := range products {
		oldPrice := product.Price * 1.2
		product.Price = product.Price * 0.9
		fmt.Printf("Товар %s теперь со скидкой! Старая цена: %.2f₽ Новая цена: %.2f₽\n", product.Title, oldPrice, product.Price)
	}
}
