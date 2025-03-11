package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5 // Уровень инфляций
	var investmentAmount float64 // Сумма инвестиций
	var expectedReturnRate float64 // Ожидаемая доходность
	var years float64 // Кол-во лет

	fmt.Print("Введите сумму инвестиций: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Ввиде кол-во лет: ")
	fmt.Scan(&years)

	fmt.Print("Ввидите ожидаемую доходность: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
