package utils

import (
	"github.com/dgrijalva/jwt-go"
	"event-ticketing-system-backend/models"
	"time"
)

var secretKey = []byte("secret")

func GenerateJWT(user models.User) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "event-ticketing-system",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
