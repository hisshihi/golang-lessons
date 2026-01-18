// Package postman provides functionalities to send and receive messages.
package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// postman симулирует работу почтальона, который доставляет письма.
//
// Параметры:
//   - ctx: Контекст для управления жизненным циклом почтальона.
//   - transferPoint: Канал для передачи доставленных писем.
//   - n: Номер почтальона.
//   - mail: Письмо, которое почтальон должен доставить.
func postman(ctx context.Context, transferPoint chan<- string, n int, mail string) {
	// Бесконечный цикл работы почтальона. Он продолжает работать, пока контекст не будет отменён.
	for {
		// Проверяем состояние контекста.
		select {
		// Если контекст отменён, почтальон завершает работу.
		case <-ctx.Done():
			fmt.Printf("Штатный почтальон номер: %d. Завершение работы\n", n)
			return
			// Если контекст не отменён, почтальон продолжает доставлять письма.
		default:
			fmt.Printf("Я почтальон номер: %d. Взял письмо: %s\n", n, mail)
			time.Sleep(1 * time.Second)
			fmt.Printf("Я почтальон номер: %d. Донёс письмо до почты: %s\n", n, mail)

			// Передаём доставленное письмо через канал.
			transferPoint <- mail
			fmt.Printf("Я почтальон номер: %d. Передал письмо: %s\n", n, mail)
		}
	}
}

// PostmanPool создаёт пул почтальонов и возвращает канал для получения доставленных писем.
//
// Параметры:
//   - ctx: Контекст для управления жизненным циклом почтальонов.
//   - postmanCount: Количество почтальонов в пуле.
//
// Возвращаемое значение:
//   - Канал для получения доставленных писем.
func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	// Создаём канал для передачи доставленных писем.
	mailTransferPoint := make(chan string)
	// Создаём группу ожидания для синхронизации завершения работы почтальонов.
	wg := &sync.WaitGroup{}

	// Запускаем указанное количество почтальонов.
	for i := 1; i <= postmanCount; i++ {
		// Используем группу ожидания для отслеживания завершения работы каждого почтальона.
		wg.Go(func() {
			postman(ctx, mailTransferPoint, i, postmanToMail(i))
		})
	}

	// Запускаем горутину для закрытия канала после завершения работы всех почтальонов.
	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

// postmanToMail возвращает письмо, соответствующее номеру почтальона.
//
// Параметры:
//   - postmanNumber: Номер почтальона.
//
// Возвращаемое значение:
//   - Строка с содержимым письма.
func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Письмо от Кожухина Дениса",
		2: "Информация о платеже",
		3: "ЖКХ квитанция",
	}
	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Новости"
	}

	return mail
}
