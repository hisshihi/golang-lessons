package main

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func hello() string {
	hello := "hello"
	return hello
}

func main() {
	hello()
}
