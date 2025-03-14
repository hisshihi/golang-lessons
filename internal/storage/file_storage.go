package storage

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var fileName = "operations.txt"

func ReadOperationsFromFile() []string {
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

func WriteOperationToFile(operation string, cash string, balance string) {
	userOperation := fmt.Sprintf("%v%v. %v. Время -> %v\n", operation, cash, balance, time.Now())

	// Открываем файл с флагами для добавления записей (O_APPEND),
	// создания если не существует (O_CREATE) и записи (O_WRONLY)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Записываем новую операцию в файл
	if _, err := file.WriteString(userOperation); err != nil {
		fmt.Println("Не удалось записать в файл", err)
	}
}
