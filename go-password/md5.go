package main

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	encodePWD := string(hash)
	return encodePWD
}

func ComparePassword(encodePWD string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePWD), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
