package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	min         float64 = 0.00
	max         float64 = 100000.00
	userBalance         = min + rand.Float64()*(max-min)
	questions   string  = "1. Проверить баланс\n2. Внести наличные\n3. Снять наличные\n4. Выйти\n5. История операций\n\n"
	fileName            = "operations.txt"
)

func readOperationsFromFile() []string {
	operations := []string{}

	// Проверяем существование файла
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("История операций пуста")
		return operations
	}

	// Читаем содержимое файла
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return operations
	}

	// Разбиваем содержимое на строки
	operations = strings.Split(string(data), "\n")

	// Удаляем пустую строку в конце, если она есть
	if operations[len(operations)-1] == "" {
		operations = operations[:len(operations)-1]
	}

	return operations
}

func writeOperationToFile(operation string, cash string, balance string) {
	userOperation := fmt.Sprintf("%v%v. %v. Время -> %v\n", operation, cash, balance, time.Now())

	// Открываем файл с флагами для добавления записей (O_APPEND),
	// создания если не существует (O_CREATE) и записи (O_WRONLY)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Записываем новую операцию в файл
	if _, err := file.WriteString(userOperation); err != nil {
		fmt.Println("Не удалось записать в файл", err)
	}
}

func main() {
	sayHello()

	var choice int
	for {
		fmt.Println(questions)
		fmt.Print("Выберите вариант из предложенных: ")
		fmt.Scan(&choice)

		if choice == 4 {
			fmt.Println("До свидания!")
			break
		}
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
		p.Printf("Ваш баланс: %.2f₽ \n", userBalance)
	case 2:
		fmt.Print("Сколько вы хотите внести? ")
		fmt.Scan(&cash)
		if cash <= 0 {
			fmt.Println("Неверная сумма для пополнения")
		} else {
			userBalance += cash
			p.Printf("Ваш баланс: %.2f₽\n", userBalance)
			writeOperationToFile("Пополнение счёта -> ", p.Sprintf("%.2f₽", cash), p.Sprintf("Текущий баланс -> %.2f₽", userBalance))
		}
	case 3:
		fmt.Print("Сколько вы хотите снять? ")
		fmt.Scan(&cash)
		if cash <= 0 {
			fmt.Println("Неверная сумма для снятия")
		} else if userBalance >= cash {
			userBalance -= cash
			p.Printf("Ваш баланс: %.2f₽\n", userBalance)
			writeOperationToFile("Снятие со счёта -> ", p.Sprintf("%.2f₽", cash), p.Sprintf("Текущий баланс -> %.2f₽", userBalance))
		} else {
			fmt.Println("Недостаточно средств")
		}
	case 5:
		fmt.Println("\nИстория операций:")
		fmt.Println("------------------------")
		operations := readOperationsFromFile()
		for _, operation := range operations {
			fmt.Println(operation)
		}
		fmt.Println("------------------------")
	default:
		fmt.Println("Неверный ввод")
	}
}
