package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(rawpass []byte) (hash string, err error) {
	hashBytes, err := bcrypt.GenerateFromPassword(rawpass, bcrypt.DefaultCost)
	hash = string(hashBytes)
	return hash, err
}

func CompareHashPassword(rawpass []byte, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), rawpass)
}

func GenerateJWTToken(userId string) (tokenstr string) {
	// var user models.User

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		// "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	secret := os.Getenv("JWT_SCRET_KEY_USER")
	tokenstr, err := token.SignedString([]byte(secret))

	if err != nil {
		panic(err)
	}
	return
}
