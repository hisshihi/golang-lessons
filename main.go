package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer func() {
		p := recover()
		if p != nil {
			err := WritePanicMsgToFile(fmt.Sprintf("Datetime: %v Panic msg: %v\n", time.Now(), p))
			if err != nil {
				fmt.Println("Error writing panic message to file:", err)
			}
		}
	}()
	slice := []int{1, 2, 3}
	fmt.Println(slice[4])
}

func WritePanicMsgToFile(panicInfo string) error {
	file, err := os.OpenFile("panic.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening or creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(panicInfo)
	if err != nil {
		return fmt.Errorf("error writing panic message to file: %v", err)
	}

	return nil
}
