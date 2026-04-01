package main

import "fmt"

type User struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func main() {
	user := User{
		ID:   1,
		Name: "hiss",
	}
	fmt.Printf("user: %#v", user)
}
