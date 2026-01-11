package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct {
}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Pay(usd int) int {
	fmt.Println("Оплата банковской картой")
	fmt.Printf("Сумма: %d USD\n", usd)

	return rand.Int()
}

func (b Bank) Cancel(id int) {
	fmt.Printf("Отмена оплаты банковской картой с ID: %d\n", id)
}
