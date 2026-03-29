package main

import "log"

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func hello() string {
	hello := "hello"
	user := User{ID: 1, Name: "hiss"}
	log.Println(user)
	return hello
}

func main() {
	hello()
}
