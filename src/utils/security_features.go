package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"html"
	"io"
	"log"
	"net/mail"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(rawpass []byte) (hash string, err error) {
	hashBytes, err := bcrypt.GenerateFromPassword(rawpass, bcrypt.DefaultCost)
	hash = string(hashBytes)
	return hash, err
}

func CompareHashPassword(rawpass, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawpass))
}

func GenerateCSRFToken(len int) (token string) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate CSRF token: %v", err)
	}
	return base64.URLEncoding.Strict().EncodeToString(bytes)
}

func GenerateJWTToken(userId string) (tokenstr string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	secret := os.Getenv("JWT_SCRET_KEY_USER")
	tokenstr, err := token.SignedString([]byte(secret))

	if err != nil {
		panic(err)
	}
	return
}

func ValidateJWTToken(tokenstr string) (exp float64, userId string, err error) {
	secret := os.Getenv("JWT_SCRET_KEY_USER")
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, err
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", errors.New("invalid token")
	}
	exp, ok = claims["exp"].(float64)
	if !ok {
		return 0, "", errors.New("invalid exp claim")
	}
	userId, ok = claims["userId"].(string)
	if !ok {
		return 0, "", errors.New("invalid userId claim")
	}

	return exp, userId, nil
}

func CalculateHash(file io.Reader) (hash string, err error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	hashBytes := hasher.Sum(nil)
	return string(hashBytes), nil

}

func SanitizeString(input string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(input)
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func EscapeString(input string) string {
	return html.EscapeString(input)
}
