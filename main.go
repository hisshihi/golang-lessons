package main

import (
	"github.com/hisshihi/golang-lessons/payments"
	"github.com/hisshihi/golang-lessons/payments/methods"
	"github.com/k0kubun/pp/v3"
)

func main() {
	method := methods.NewCrypto()

	paymentModule := payments.NewPaymentModule(method)

	operation1 := paymentModule.Pay("Шоколадная колбаска", 2)
	paymentModule.Pay("Кофе", 3)
	paymentModule.Pay("Сендвич", 4)

	allInfo := paymentModule.AllInfo()
	pp.Println("Все операции", allInfo)
	
	info := paymentModule.Info(operation1)
	pp.Println("Информация по операции 1", info)
}
