package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(password string) string {
	hashResult, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("加密失败")
		return ""
	}
	encodedPassword := string(hashResult)
	return encodedPassword
}

func PasswordCompare(password string, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
