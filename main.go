package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// writer генерирует числа от 1 до 10 и отправляет их в канал
func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

func doubler(in <-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	for num := range in {
		wg.Add(1)
		go func(n int) {
			time.Sleep(500 * time.Millisecond)
			out <- n * 2
			wg.Done()
		}(num)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func reader(in <-chan int) {
	var results []int
	for num := range in {
		results = append(results, num)
	}
	sort.Ints(results)
	for _, num := range results {
		fmt.Println(num)
	}
}

func main() {
	reader(doubler(writer()))
}
