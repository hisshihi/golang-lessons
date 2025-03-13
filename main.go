package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	min         float64 = 0.00
	max         float64 = 100000.00
	userBalance         = min + rand.Float64()*(max-min)
	questions   string  = "1. Проверить баланс\n2. Внести наличные\n3. Снять наличные\n4. Выйти\n\n"
)

func main() {
	sayHello()

	var choice int
	for choice != 4 {
		fmt.Println(questions)
		fmt.Print("Выберите вариант из предложенных: ")
		fmt.Scan(&choice)

		userChoice(choice)
	}
}

func sayHello() {
	now := time.Now().Hour()
	if now >= 23 && now <= 6 {
		fmt.Println("Доброй ночи!")
	} else if now > 6 && now <= 11 {
		fmt.Println("Доброе утро!")
	} else if now > 11 && now <= 16 {
		fmt.Println("Добрый день!")
	} else {
		fmt.Println("Добрый вечер!")
	}
}

func userChoice(choice int) {
	var cash float64
	p := message.NewPrinter(language.Russian)

	switch choice {
	case 1:
		p.Printf("Ваш баланс: %.2f \n", userBalance)
	case 2:
		fmt.Print("Сколько вы хотите внести? ")
		fmt.Scan(&cash)
		userBalance += cash
		p.Printf("Ваш баланс: %.2f\n", userBalance)
	case 3:
		fmt.Print("Сколько вы хотите снять? ")
		fmt.Scan(&cash)
		userBalance -= cash
		p.Printf("Ваш баланс: %.2f\n", userBalance)
	default:
		fmt.Println("Неверный ввод")
	}
}
