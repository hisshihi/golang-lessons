package main

import (
	"github.com/hisshihi/golang-lessons/payments"
	"github.com/hisshihi/golang-lessons/payments/methods"
	"github.com/k0kubun/pp/v3"
)

func main() {
	method := methods.NewBonus()

	paymentMethod := payments.NewPaymentModule(method)

	operation1 := paymentMethod.Pay("Шоколадная колбаска", 2)
	paymentMethod.Pay("Кофе", 3)
	paymentMethod.Pay("Сендвич", 4)

	allInfo := paymentMethod.AllInfo()
	pp.Println("Все операции", allInfo)
	
	paymentMethod.Cancel(operation1)
	
	info := paymentMethod.Info(operation1)
	pp.Println("Информация по операции", info)
}
