package models

import "github.com/hisshihi/golang-lessons/db"

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

	result, err := stmt.Exec(u.Email, u.Password)
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
