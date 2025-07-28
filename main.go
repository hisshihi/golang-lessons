package main

import "fmt"

type Product struct {
	ID    int
	Title string
	Price float64
}

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites["Amazon Web Services"])
	fmt.Println(websites)
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
