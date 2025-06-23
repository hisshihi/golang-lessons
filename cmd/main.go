package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hisshihi/golang-lessons/note"
)

type Saver interface {
	Save() error
}

type outputtable interface {
	Saver
	Display()
}

func main() {
	// printSomething(1)
	// printSomething(1.5)
	// printSomething("Hello, World!")

	fmt.Println("Welcome to the Note Taking App!")
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	outputtable(userNote).Display()
}

// printSomething может принимать разные типы данных и выводить их на экран
func printSomething(value any) {
	// typeVal, ok := value.(int)
	// if ok {
	// 	fmt.Println("Integer value:", typeVal)
	// } else {
	// 	fmt.Println("Value is not an integer, it is of type:", fmt.Sprintf("%T", value))
	// }
	// switch v := value.(type) {
	// case string:
	// 	fmt.Println("String value:", v)
	// case int:
	// 	fmt.Println("Integer value:", v)
	// default:
	// 	fmt.Println("Unknown type")
	// }
}


func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	readString = strings.TrimSuffix(readString, "\n")
	readString = strings.TrimSuffix(readString, "\r")

	return readString
}
