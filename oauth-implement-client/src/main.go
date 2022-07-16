package main

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

var mySignedKey = []byte("sampleSignedKey")

func GenerateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "@binator_1308"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySignedKey)

	if err != nil {
		fmt.Errorf("generating JWT Token failed")
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("Hello World")

	tokenString, err := GenerateJWTToken()
	if err != nil {
		fmt.Println("Error in generating token")
	}

	fmt.Println(tokenString)
}
