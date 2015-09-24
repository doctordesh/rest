package main

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func authUserPassword(user *User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err == nil {
		return true
	}
	return false
}

func authUser(email string, password string) (*User, error) {
	var user *User
	for _, tmp := range users {
		if tmp.Email == email {
			user = &tmp
		}
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	if authUserPassword(user, password) == false {
		return nil, errors.New("User auth failed")
	}

	return user, nil
}
