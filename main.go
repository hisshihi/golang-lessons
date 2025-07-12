package main

import "fmt"

func main() {
	// 1 lesson
	hobbies := [3]string{"coding", "gaming", "learning"}
	for i, hobby := range hobbies {
		fmt.Printf("My hobby #%d: %s\n", i+1, hobby)
	}

	// 2 lesson
	firstHobby := hobbies[0]
	fmt.Printf("My first hobby: %s\n", firstHobby)
	otherHobbies := hobbies[1:]
	fmt.Printf("Other hobbies: %v\n", otherHobbies)

	// 3 lesson
	firstAndSecondHobbies := hobbies[:2]
	fmt.Printf("First and second hobbies: %v\n", firstAndSecondHobbies)
	firstAndSecondHobbiesNewSlice := []string{}
	firstAndSecondHobbiesNewSlice = append(firstAndSecondHobbiesNewSlice, hobbies[:2]...)
	fmt.Printf("First and second hobbies new slice: %v\n", firstAndSecondHobbiesNewSlice)

	// 4 lesson
	firstAndSecondHobbies = hobbies[1:]
	fmt.Printf("First and second hobbies after modification: %v\n", firstAndSecondHobbies)

	// 5 lesson
	todo := []string{"Learn Go", "Practice coding"}
	fmt.Printf("TODO: %v\n", todo)

	// 6 lesson
	todo[1] = "Practice Go"
	todo = append(todo, "Build a project")
	fmt.Printf("Updated TODO: %v\n", todo)

	// 7 lesson
	type Product struct {
		Title string
		ID int64
		Price float64
	}

	products := []Product{
		{Title: "Laptop", ID: 1, Price: 999.99},
		{Title: "Smartphone", ID: 2, Price: 499.99},
	}

	products = append(products, Product{Title: "MacBook Pro", ID: 3, Price: 1299.99})
	fmt.Printf("Products: %v\n", products)
}
