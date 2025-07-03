package main

import "fmt"

func main() {
	fmt.Println(add(32, 22))
}

func add(a, b any) any {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}
	return nil
}
