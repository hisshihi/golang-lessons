package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 0, 3)

	userNames = append(userNames, "Denis")
	userNames = append(userNames, "Arina")
	userNames = append(userNames, "Musinia")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 2)

	courseRatings["go"] = 5.0
	courseRatings["java"] = 4.8

	courseRatings.output()
}
