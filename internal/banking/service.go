package banking

import (
	"fmt"
	"math/rand/v2"

	"github.com/hisshihi/bank/internal/storage"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	min         float64 = 0.00
	max         float64 = 100000.00
	userBalance         = min + rand.Float64()*(max-min)
)

func UserChoice(choice int) {
	var cash float64
	p := message.NewPrinter(language.Russian)

	switch choice {
	case 1:
		p.Printf("Ваш баланс: %.2f₽ \n", userBalance)
	case 2:
		fmt.Print("Сколько вы хотите внести? ")
		fmt.Scan(&cash)
		if cash <= 0 {
			fmt.Println("Неверная сумма для пополнения")
		} else {
			Deposit(cash, userBalance)
		}
	case 3:
		fmt.Print("Сколько вы хотите снять? ")
		fmt.Scan(&cash)
		if cash <= 0 {
			fmt.Println("Неверная сумма для снятия")
		} else if userBalance >= cash {
			Withdraw(cash, userBalance)
		} else {
			fmt.Println("Недостаточно средств")
		}
	case 5:
		fmt.Println("\nИстория операций:")
		fmt.Println("------------------------")
		operations := storage.ReadOperationsFromFile()
		for _, operation := range operations {
			fmt.Println(operation)
		}
		fmt.Println("------------------------")
	default:
		fmt.Println("Неверный ввод")
	}
}
