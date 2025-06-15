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

func main() {
	fmt.Println("Welcome to the Note Taking App!")
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}

	fmt.Println("Data saved successfully!")
	return nil
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
