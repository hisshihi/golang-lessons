package banking

import (
	"github.com/hisshihi/golang-lessons/internal/storage"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var p = message.NewPrinter(language.Russian)

func Deposit(cash float64, userBalance float64) {
	userBalance += cash
	p.Printf("Ваш баланс: %.2f₽\n", userBalance)
	storage.WriteOperationToFile("Пополнение счёта -> ", p.Sprintf("%.2f₽", cash), p.Sprintf("Текущий баланс -> %.2f₽", userBalance))
}

func Withdraw(cash float64, userBalance float64) {
	userBalance -= cash
	p.Printf("Ваш баланс: %.2f₽\n", userBalance)
	storage.WriteOperationToFile("Снятие со счёта -> ", p.Sprintf("%.2f₽", cash), p.Sprintf("Текущий баланс -> %.2f₽", userBalance))
}
