package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(text string, countPostman int) {
	for i := range countPostman {
		fmt.Printf("Я почтальён. Я отнёс газету %s в %d раз\n", text, i+1)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	namesNews := []string{"Новости", "IT и технологии", "Игровые новости"}

	for _, name := range namesNews {
		wg.Go(func() {
			postman(name, len(namesNews))
		})
	}
	wg.Wait()
}
