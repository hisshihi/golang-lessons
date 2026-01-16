package main

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0

func setLike(mu *sync.RWMutex) {
	for range 100_000 {
		mu.Lock()
		likes++
		mu.Unlock()
	}
}

func getLike(mu *sync.RWMutex) {
	for range 100_000 {
		mu.RLock()
		_ = likes
		mu.RUnlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	rwMu := &sync.RWMutex{}

	initTime := time.Now()

	for range 10 {
		wg.Go(func() {
			setLike(rwMu)
		})
	}

	for range 10 {
		wg.Go(func() {
			getLike(rwMu)
		})
	}

	wg.Wait()

	fmt.Println("Time taken:", time.Since(initTime))
}
