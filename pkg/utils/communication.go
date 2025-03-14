package utils

import (
	"fmt"
	"time"

	"github.com/Pallinder/go-randomdata"
)

func SayHello() {
	randomName := randomdata.SillyName()
	now := time.Now().Hour()
	if now >= 23 && now <= 6 {
		fmt.Println("Доброй ночи!", randomName)
	} else if now > 6 && now <= 11 {
		fmt.Println("Доброе утро!", randomName)
	} else if now > 11 && now <= 16 {
		fmt.Println("Добрый день!", randomName)
	} else {
		fmt.Println("Добрый вечер!", randomName)
	}
}
