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
		for i := range 10_000 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

// doubler принимает числа из входного канала, удваивает их с задержкой и отправляет в выходной канал
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

// reader принимает числа из входного канала, сортирует их и выводит на экран
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
	now := time.Now()
	reader(doubler(writer()))
	defer func() {
		fmt.Println(time.Since(now))
	}()
}
