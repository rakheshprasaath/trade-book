package models

import "golang.org/x/crypto/bcrypt"

type User struct{
	Id int `json: "id"`
	UserName string `json: "user_name"`
	Email string `json: "email"`
	Phone string `json: "phone"`
	Password []byte `json: "password"`
}

// SetPassword function to hash the user's password
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}