package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() {
	randInt := rand.Intn(10) + 1
	fmt.Println(randInt)
	time.Sleep(time.Duration(randInt) * time.Second)
}

// predictbleTimeWork функция обёртка, которая будет прервывать выполнение если функция randomTimeWork работает дольше 3 секунд
func predictbleTimeWork() {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
	case <-time.After(3 * time.Second):
	}
}

func main() {
	initTime := time.Now()
	predictbleTimeWork()
	fmt.Println(time.Since(initTime))
}
