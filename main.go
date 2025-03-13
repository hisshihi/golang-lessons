package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Доход: ")
	expenses := getUserInput("Расходы: ")
	taxRate := getUserInput("Налоговая ставка: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit

	return
}

func getUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}
