package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct {}

func NewPayPal() PayPal {
	return PayPal{}
}

func (p PayPal) Pay(usd int) int {
	fmt.Println("Оплата через PayPal")
	fmt.Printf("Сумма: %d USD\n", usd)

	return rand.Int()
}

func (p PayPal) Cancel(id int) {
	fmt.Printf("Отмена оплаты через PayPal с ID: %d\n", id)
}