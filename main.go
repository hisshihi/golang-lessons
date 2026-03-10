package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chanForResp := make(chan resp)
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	go RPCCall(ctxWithTimeout, chanForResp)
	res:= <-chanForResp
	fmt.Println(res.id, res.err)
}

type resp struct {
	id  int
	err error
}

func RPCCall(ctx context.Context, ch chan<- resp) {
	duration := rand.Intn(5) + 1
	fmt.Println("duration:", duration)

	select {
	case <-ctx.Done():
		ch <- resp{
			id:  0,
			err: errors.New("timeout expired"),
		}
		return
	case <-time.After(time.Second * time.Duration(duration)):
		ch <- resp{
			id:  rand.Int(),
			err: nil,
		}
		return
	}
}
