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
