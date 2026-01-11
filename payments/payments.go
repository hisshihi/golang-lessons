// Package payments реализует модуль оплаты
package payments

import "maps"

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentInfo   map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentInfo:   make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Pay совершить оплату
//
// Принимает:
// 1. Описание проводимой оплаты
// 2. Сумма оплаты
// Возвращает:
// 1. ID проведённой операции.
func (m *PaymentModule) Pay(description string, usd int) int {
	// 1. Проводим операцию
	// 2. Получаем id операции
	id := m.paymentMethod.Pay(usd)

	// 3. Сохранеям данные об операции
	// - описание операции
	// - сколько было потрачено
	// - отменялась ли операция
	info := PaymentInfo{
		Description: description,
		CountUSD:    usd,
		IsCancelled: false,
	}
	m.paymentInfo[id] = info

	//	4. Возвращать айди проведённой операции
	return id
}

// Cancel отменить оплату
//
// Принимает:
// 1. ID опреации
// Возвращает:
// - ничего не возвращает.
func (m *PaymentModule) Cancel(id int) {
	info, exists := m.paymentInfo[id]
	if !exists {
		return
	}

	m.paymentMethod.Cancel(id)

	info.IsCancelled = true
	m.paymentInfo[id] = info
}

// Info информация о конкретной операции
//
// Принимает:
// 1. ID операции
// Возвращает:
// - возврщает информацию о проведённой операции.
func (m *PaymentModule) Info(id int) PaymentInfo {
	info, exists := m.paymentInfo[id]
	if !exists {
		return PaymentInfo{}
	}
	return info
}

// AllInfo информация о всех операция
//
// Принимает:
// - ничего не принимает
// Возвращает:
// 1. Все проведённые операции.
func (m *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(m.paymentInfo))
	maps.Copy(tempMap, m.paymentInfo)
	return tempMap
}
