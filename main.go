package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func rpcCall(ctx context.Context) (string, error) {
	duration := time.Duration(rand.Intn(10)+1) * time.Second
	fmt.Println("duration", duration)
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-timer.C:
		return "data", nil
	case <-ctx.Done():
		return "", fmt.Errorf("rpc error: %w", ctx.Err())
	}
}

func callRPCWithTimeout(ctx context.Context, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	type result struct {
		data string
		err  error
	}

	ch := make(chan result, 1)
	go func() {
		data, err := rpcCall(ctx)
		ch <- result{data: data, err: err}
	}()

	select {
	case res := <-ch:
		return res.data, res.err
	case <-ctx.Done():
		return "", fmt.Errorf("timeout: %w", ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	data, err := callRPCWithTimeout(ctx, 5*time.Second)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println(data)
}
