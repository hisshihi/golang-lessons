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
	
	fmt.Print("Ввидите ожидаемую доходность: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Ввиде кол-во лет: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}
