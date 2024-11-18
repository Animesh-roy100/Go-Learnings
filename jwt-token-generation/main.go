package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	secret := []byte("your-secret-key")
	claims := jwt.RegisteredClaims{
		Subject:   "user123",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), // 1-hour expiration
		Issuer:    "api-gateway",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated Token:", signedToken)
}
