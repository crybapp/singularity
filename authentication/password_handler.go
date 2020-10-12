package authentication

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

const passwordCost = 8

//EncryptPassword will output a hash of the current password, to be stored.
func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}

//VerifyPassword compares the hash and the input string and checks if they're the same.
func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
