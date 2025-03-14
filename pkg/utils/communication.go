package utils

import (
	"fmt"
	"time"
)

func SayHello() {
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
