package models

import (
	"github.com/hisshihi/golang-lessons/db"
	"github.com/hisshihi/golang-lessons/utils"
)

const (
	ErrUserNotFound = "sql: no rows in result set"
	ErrUniqueEmail  = "UNIQUE constraint failed: users.email"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *User) Save() error {
	q := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	HashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, HashPassword)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(lastID)

	return nil
}

func GetUserByEmail(email string) (*User, error) {
	q := `
		SELECT email, password
		FROM users
		WHERE email = ?
	`
	var user User
	err := db.DB.QueryRow(q, email).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
