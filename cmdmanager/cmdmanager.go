// Package cmdmanager работа с терминалом
package cmdmanager

import (
	"fmt"
	"strconv"
)

type CMDManager struct{}

// ReadLines ввод данных от пользователя в терминал
func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Введите ваши цены. После чего нажмите Enter.")
	fmt.Println("Для выполнения расчёта введите любой символ.")

	prices := []string{}

Loop:
	for {
		var price string
		fmt.Print("Цена: ")
		fmt.Scan(&price)
		priceInt, err := strconv.ParseFloat(price, 64)
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok {
				switch numErr.Err {
				case strconv.ErrSyntax:
					break Loop
				default:
					fmt.Println("Возникла ошибка ввода, попробуйте ещё раз.")
					continue
				}
			}
		}

		if priceInt <= 0 {
			fmt.Println("Число не может быть меньше или равно нулю")
			continue
		}
		prices = append(prices, price)
	}

	return prices, nil
}

// WriteResult вывод данных в терминал
func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func NewCMDManager() CMDManager {
	return CMDManager{}
}
