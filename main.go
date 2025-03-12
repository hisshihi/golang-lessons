package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5 // Уровень инфляций
var (
	investmentAmount   float64 = 1000.0 // Сумма инвестиций
	expectedReturnRate float64 = 5.5 // Ожидаемая доходность
	years              float64 = 10.0 // Кол-во лет
)

func main() {
	askAQuestion()

	futureValue, futureRealValue := calculateFutureValue()

	formattedFV := fmt.Sprintf("Future value %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func calculateFutureValue() (fv, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return
}

func askAQuestion() {
	fmt.Print("Введите сумму инвестиций: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Ввидите ожидаемую доходность: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Ввиде кол-во лет: ")
	fmt.Scan(&years)
}
