package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type outVal struct {
	val int
	err error
}

var errTimeout = errors.New("timed out")

func processData(ctx context.Context, val int) chan outVal {
	ch := make(chan int)
	out := make(chan outVal)

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(ch)
	}()

	go func() {
		select {
		case <-ch:
			out <- outVal{
				val: val * 2,
				err: nil,
			}
		case <-ctx.Done():
			out <- outVal{
				val: 0,
				err: errTimeout,
			}
		}
	}()

	return out
}

func main() {
	inCh := make(chan int)
	outCh := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		defer close(inCh)
		for i := range 10 {
			select {
			case inCh <- i + 1:
			case <-ctx.Done():
				return
			}
		}
	}()

	now := time.Now()
	processParallel(ctx, inCh, outCh, 5)

	for val := range outCh {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}

func processParallel(ctx context.Context, inCh <-chan int, outCh chan<- int, workerCount int) {
	wg := &sync.WaitGroup{}
	for range workerCount {
		wg.Add(1)
		go worker(ctx, inCh, outCh, wg)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()
}

func worker(ctx context.Context, inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case val, ok := <-inCh:
			if !ok {
				return
			}
			select {
			case ov := <-processData(ctx, val):
				if ov.err != nil {
					return
				}
				select {
				case outCh <- ov.val:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
