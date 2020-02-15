package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey []byte

func init() {
	secretKey = []byte("SecretKeyDate")
}

func main() {
	claims := jwt.MapClaims{
		"id":  1,
		"nbf": time.Now(),
	}

	token, err := issue(claims)
	if err != nil {
		fmt.Println(err)
	}

	_claims, err := verify(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("id %v\n", _claims["id"])
	fmt.Printf("nbf %v\n", _claims["nbf"])
}

func issue(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verify(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
