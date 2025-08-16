// Package filemanager работа с файлами
package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// Загружаем файл и получаем данные о ценах
	file, err := os.Open(path)
	if err != nil {
		return []string{}, fmt.Errorf("возникла ошибка с чтением файла: %v", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return []string{}, fmt.Errorf("возникла ошибка с чтением данных из файла: %v", err.Error())
	}

	return lines, nil
}

func WriteJSON(data any, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("не удалось преобразовать данные в JSON: %v", err)
	}

	defer file.Close()
	return nil
}
