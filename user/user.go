package user

import (
	"errors"
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

type Admin struct {
	User     User
	email    string
	password string
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "---",
			birthDate: "---",
			createdAt: time.Now(),
			id:        uuid.NewString(),
		},
	}, nil
}

func New(firstName, lastName, birthDate string) (*User, error) {
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
