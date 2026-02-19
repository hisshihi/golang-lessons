package main

import (
	"fmt"
	"math/rand"
	"time"
)

var maxWaitSeconds = 5

// randomWait вызов долго работающей функции
func randomWait() int {
	workSeconds := rand.Intn(maxWaitSeconds + 1)

	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

func main() {
	totalWorkSeconds := 0
	now := time.Now()

	ch := make(chan int)

	for range 100 {
		go func() {
			workSeconds := randomWait()
			ch <- workSeconds
		}()
	}
 
	for range 100 {
		totalWorkSeconds += <- ch
	}

	fmt.Println(time.Since(now))
	fmt.Println("work seconds:", totalWorkSeconds)
}
