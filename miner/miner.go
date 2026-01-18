// Package miner provides functionality for simulating a miner that extracts and transfers coal.
package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// miner симулирует работу шахтёра, который добывает уголь.
//
// Параметры:
//   - ctx: Контекст для управления жизненным циклом шахтёра.
//   - transferPoint: Канал для передачи добытого угля.
//   - n: Номер шахтёра.
//   - power: Количество угля, которое шахтёр добывает за цикл.
func miner(
	ctx context.Context,
	transferPoint chan<- int,
	n int,
	power int,
) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Шахтёр номер: %d. Завершение работы\n", n)
			return
		default:
			fmt.Printf("Я шахтёр номер: %d. Начал добывать уголь\n", n)
			time.Sleep(1 * time.Second)
			fmt.Printf("Я шартёр номер: %d. Добыл уголь: %d\n", n, power)

			transferPoint <- power
			fmt.Printf("Я шахтёр номер: %d. Передал уголь: %d\n", n, power)
		}
	}
}

// MinerPool создаёт пул шахтёров и возвращает канал для получения добытого угля.
//
// Параметры:
//   - ctx: Контекст для управления жизненным циклом шахтёров.
//   - minerCount: Количество шахтёров в пуле.
//
// Возвращаемое значение:
//   - Канал для получения добытого угля.
func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Go(func() {
			miner(ctx, coalTransferPoint, i, i*10)
		})
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
