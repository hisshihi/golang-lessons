package user

import "fmt"

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
