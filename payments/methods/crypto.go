// Package methods предоставляет различные методы оплаты.
package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct {}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Оплата криптовалютой")
	fmt.Printf("Сумма: %d USD\n", usd)

	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	fmt.Printf("Отмена оплаты криптовалютой с ID: %d\n", id)
}