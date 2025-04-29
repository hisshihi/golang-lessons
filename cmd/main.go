package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	appUser := &user{
		firstName: getUserString("Enter your first name:"),
		lastName:  getUserString("Enter your last name:"),
		birthDate: getUserString("Enter your birth date (YYYY-MM-DD):"),
		createdAt: time.Now(),
	}

	appUser.outputUserDetails()
}

// Receiver
func (u *user) outputUserDetails() {
	fmt.Printf("First Name: %s\n", u.firstName)
	fmt.Printf("Last Name: %s\n", u.lastName)
	fmt.Printf("Birth Date: %s\n", u.birthDate)
	fmt.Printf("User created at %v\n", u.createdAt.Format("2006-01-02 15:04:05"))
}

func getUserString(promt string) string {
	fmt.Println(promt)
	var input string
	fmt.Scanln(&input)

	return input
}
