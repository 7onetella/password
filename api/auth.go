package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// IDClaims custom jwt claim with ID
type IDClaims struct {
	ID string
	jwt.StandardClaims
}

// EncodeID signs id
func EncodeID(id string) (string, time.Time, error) {
	expTime := time.Now().Add(20 * time.Second)
	claim := IDClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)

	return tokenString, expTime, err
}

// DecodeToken decodes jwt token
func DecodeToken(tokenString string) (string, int64, error) {

	token, err := jwt.ParseWithClaims(tokenString, &IDClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return "", 0, err
	}

	if claims, ok := token.Claims.(*IDClaims); ok && token.Valid {
		return claims.ID, claims.ExpiresAt, nil
	}

	return "", 0, errors.New("Token Invalid")

}
