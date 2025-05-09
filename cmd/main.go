package main

import (
	"fmt"

	"github.com/hisshihi/golang-lessons/user"
)

func main() {
	fmt.Println("Welcome to the User Management System")
	var appUser user.User

	admin, err := user.NewAdmin("admin@example.com", "admin")
	if err != nil {
		fmt.Println("Error creating admin:", err)
	}

	admin.User.SaveAdmin()

	for {
		fmt.Println("1. Add User")
		fmt.Println("2. Get User")
		fmt.Println("3. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			firstName := userInput("Enter first name:")
			lastName := userInput("Enter last name:")
			birthDate := userInput("Enter birth date (DD.MM.YYYY):")
			appUser, err := user.New(firstName, lastName, birthDate)
			if err != nil {
				fmt.Println("Error creating user:", err)
			}

			appUser.SaveUser()
		case 2:
			fmt.Println("Enter first name to get user details:")
			var chosenName string
			fmt.Scanln(&chosenName)
			fmt.Println(appUser.GetUser(chosenName))
		case 3:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func userInput(promt string) string {
	fmt.Println(promt)
	var input string
	fmt.Scanln(&input)

	return input
}
