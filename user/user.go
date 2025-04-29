package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        string
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		id:        uuid.NewString(),
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

var database = make(map[string]User)

// Receiver
func (u *User) SaveUser() {
	database[u.firstName] = *u
}

func (u *User) GetUser(firstName string) string {
	user, exists := database[firstName]
	if !exists {
		fmt.Println("User not found")
	}
	return fmt.Sprintf("User: %s %s %s, Birth Date: %s, Created At: %s", user.id, user.firstName, user.lastName, user.birthDate, user.createdAt)
}
