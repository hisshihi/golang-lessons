package main

import "fmt"

type str string

func (s str) log() {
	fmt.Println("Logging:", s)
}

func main() {
	var name str = "Denis"

	name.log()
}
