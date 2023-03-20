package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string
}

func NewUser(name, email, password string) *User {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return &User{
		Name:     name,
		Email:    email,
		password: string(hashPassword),
	}
}
