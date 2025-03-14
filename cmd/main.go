package main

import (
	"fmt"

	"github.com/hisshihi/bank/internal/banking"
	"github.com/hisshihi/bank/pkg/utils"
)

var questions string = "1. Проверить баланс\n2. Внести наличные\n3. Снять наличные\n4. Выйти\n5. История операций\n\n"

func main() {
	utils.SayHello()

	var choice int
	for {
		fmt.Println(questions)
		fmt.Print("Выберите вариант из предложенных: ")
		fmt.Scan(&choice)

		if choice == 4 {
			fmt.Println("До свидания!")
			break
		}
		banking.UserChoice(choice)
	}
}
