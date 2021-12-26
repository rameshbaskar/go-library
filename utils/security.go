package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(plainText string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	return string(hash)
}

func IsPasswordValid(encPassword string, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encPassword), []byte(plainText))
	if err != nil {
		return false
	}
	return true
}