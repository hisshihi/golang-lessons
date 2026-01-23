package main

import (
	"fmt"
	"math/rand"
	"time"
)

// randWait симулирует случайную работу от 1 до 5 секунд и возвращает количество секунд работы.
func randWait() int {
	workSeconds := rand.Intn(5) + 1 // Случайное число от 1 до 5

	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

func main() {
	ch := make(chan int)
	totalWorkSeconds := 0

	initTime := time.Now()

	for range 100 {
		go func() {
			seconds := randWait()

			ch <- seconds
		}()
	}

	for range 100 {
		totalWorkSeconds += <-ch
	}

	mainSecond := time.Since(initTime)
	fmt.Println("main", mainSecond)
	fmt.Println("total", totalWorkSeconds, "seconds")
}
