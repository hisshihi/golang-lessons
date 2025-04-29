package main

import (
	"fmt"
	"time"
)

type user struct {
	id        int
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	fmt.Println("Welcome to the User Management System")
	var appUser *user
	for {
		fmt.Println("1. Add User")
		fmt.Println("2. Get User")
		fmt.Println("3. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			appUser = &user{
				firstName: getUserString("Enter first name:"),
				lastName:  getUserString("Enter last name:"),
				birthDate: getUserString("Enter birth date (DD.MM.YYYY):"),
				createdAt: time.Now(),
			}
			appUser.saveUser()
		case 2:
			fmt.Println("Enter first name to get user details:")
			var chosenName string
			fmt.Scanln(&chosenName)
			fmt.Println(appUser.getUser(chosenName))
		case 3:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

var database = make(map[string]user)

// Receiver
func (u *user) saveUser() {
	id := 0
	id += 1
	u.id = id
	database[u.firstName] = *u
}

func (u *user) getUser(firstName string) string {
	user, exists := database[firstName]
	if !exists {
		fmt.Println("User not found")
	}
	return fmt.Sprintf("User: %d %s %s, Birth Date: %s, Created At: %s", user.id, user.firstName, user.lastName, user.birthDate, user.createdAt)
}

func getUserString(promt string) string {
	fmt.Println(promt)
	var input string
	fmt.Scanln(&input)

	return input
}
