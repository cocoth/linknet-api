package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(rawpass []byte) (hash string, err error) {
	hashBytes, err := bcrypt.GenerateFromPassword(rawpass, bcrypt.DefaultCost)
	hash = string(hashBytes)
	return hash, err
}
