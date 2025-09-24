package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func simpleSearchInSlice[T comparable](slice []T, target T) T {
	for _, item := range slice {
		if item == target {
			return item
		}
	}
	return *new(T)
}

func parallelSearch[T comparable](slice []T, target T) T {
	numGorutines := 4
	length := len(slice)
	results := make(chan T, numGorutines)

	segmentSize := (length + numGorutines - 1) / numGorutines

	for i := range numGorutines {
		start := i * segmentSize
		end := start + segmentSize
		if end > length {
			end = length
		}

		wg.Add(1)

		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				if slice[j] == target {
					results <- slice[j]
					return
				}
			}
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return <-results
}

func main() {
	slice := make([]int, 10000000)
	for i := range slice {
		slice[i] = i * 3
	}
	now := time.Now()
	item := simpleSearchInSlice(slice, 9213456)
	fmt.Println(item)
	fmt.Println(time.Since(now))

	now = time.Now()
	item = parallelSearch(slice, 9213456)
	fmt.Println(item)
	fmt.Println(time.Since(now))
}
