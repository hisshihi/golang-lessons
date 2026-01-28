package main

import (
	"fmt"
	"math/rand"
	"time"
)

// writer генерирует числа от 1 до 10
func writer() <-chan int {
	ch := make(chan int)
	go func() {
		for range 10 {
			value := rand.Intn(10) + 2
			ch <- value
			fmt.Println("writer value - ", value)
		}
		close(ch)
	}()

	return ch
}

// double умножает числа на 2, имитируя работу (500ms)
func double(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for num := range input {
			time.Sleep(500 * time.Millisecond)
			value := num * 2
			output <- value
			fmt.Println("double value - ", value)
		}
		close(output)
	}()

	return output
}

// reader читает и выводит на экран
func reader(input <-chan int) {
	for num := range input {
		fmt.Println("output value - ", num)
	}
}

func main() {
	reader(double(writer()))
}
