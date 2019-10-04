package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret = []byte("hello")

// IDClaims custom jwt claim with ID
type IDClaims struct {
	ID string
	jwt.StandardClaims
}

// EncodeID signs id
func EncodeID(id string) (string, error) {
	claim := IDClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

// DecodeToken decodes jwt token
func DecodeToken(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &IDClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*IDClaims); ok && token.Valid {
		return claims.ID, nil
	} else {
		return "", errors.New("Token Invalid")
	}

}
