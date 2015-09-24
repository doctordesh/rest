package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	Token        string    `json:"token"`
	PasswordHash string    `json:"-"`
	Accepted     string    `json:"accepted"`
	AcceptedAt   time.Time `json:"accepted_at"`
}

type Users []User

var users Users

func init() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("asdqwe123"), bcrypt.DefaultCost)
	users = append(users, User{
		Email:        "emil@dd.com",
		Name:         "Emil Rosendahl",
		Code:         "1337",
		Token:        "aabbcc-ddeeff-gghhii-jjkkll-mmnnoo",
		PasswordHash: string(hash),
	})

	fmt.Printf("Users: %v\n", users)
}
