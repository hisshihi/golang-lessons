package methods

import (
	"fmt"
	"math/rand"
)

type Bonus struct {}

func NewBonus() Bonus {
	return Bonus{}
}

func (b Bonus) Pay(usd int) int {
	fmt.Println("Оплата бонусами")
	fmt.Printf("Сумма: %d USD\n", usd)

	return rand.Int()
}

func (b Bonus) Cancel(id int)  {
	fmt.Println("Отмена операции бонусами с ID:", id)
}