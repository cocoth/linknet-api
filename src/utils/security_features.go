package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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

func CalculateHash(filePath string) (hash string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func CalculateHashByBuffer(file io.Reader) string {
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Encrypt encrypts plaintext using AES encryption
func Encrypt(plaintext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using AES encryption
func Decrypt(ciphertext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertextBytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	return string(ciphertextBytes), nil
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
