package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenSigned, err := token.SignedString([]byte("secreto"))
	if err != nil {
		return "", err
	}
	return tokenSigned, nil
}

func ParseToken() {
}
