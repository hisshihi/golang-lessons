package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hisshihi/golang-lessons/miner"
	"github.com/hisshihi/golang-lessons/postman"
)

func main() {
	var coal atomic.Int64
	var mails []string

	ctx := context.Background()
	minerContext, minerCancel := context.WithCancel(ctx)
	postmanContext, postmanCancel := context.WithCancel(ctx)

	coalTransferPoint := miner.MinerPool(minerContext, 2)
	mailTransferPoint := postman.PostmanPool(postmanContext, 2)

	initTime := time.Now()
	fmt.Println("<<-----------------Рабочий день шахтёров и почтальонов начался----------------->>")

	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
		fmt.Println("<<-----------------Рабочий день шахтёров окончен----------------->>")
	}()

	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
		fmt.Println("<<-----------------Рабочий день почтальонов окончен----------------->>")
	}()

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Go(func() {
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	})

	wg.Go(func() {
		for mail := range mailTransferPoint {
			mu.Lock()
			mails = append(mails, mail)
			mu.Unlock()
		}
	})

	wg.Wait()

	fmt.Println("<<-----------------Подведение итогов рабочего дня----------------->>")
	fmt.Printf("Рабочий день длился: %v\n", time.Since(initTime))

	fmt.Printf("Суммарное кол-во добытого угля: %d\n", coal.Load())

	mu.Lock()
	fmt.Printf("Суммарное кол-во полученной почты: %v\n", len(mails))
	mu.Unlock()
}
