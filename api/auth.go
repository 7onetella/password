package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var hmacSecret = []byte("")

func init() {
	cryptoToken := os.Getenv("CRYPTO_TOKEN")
	if len(cryptoToken) == 0 {
		fmt.Println("CRYPTO_TOKEN is empty")
		os.Exit(1)
	}
}

// IDClaims custom jwt claim with ID
type IDClaims struct {
	ID string
	jwt.StandardClaims
}

// EncodeID signs id
func EncodeID(id string) (string, time.Time, error) {
	expiration := time.Now().Add(5 * time.Minute)
	claim := IDClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)

	return tokenString, expiration, err
}

// DecodeToken decodes jwt token
func DecodeToken(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &IDClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*IDClaims); ok && token.Valid {
		return claims.ID, nil
	}

	return "", errors.New("Token Invalid")

}
