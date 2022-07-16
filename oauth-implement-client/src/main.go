package main

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

// Should be picked from env variables or config
var mySignedKey = []byte("sampleSignedKey")

func HomePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWTToken()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

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

func handleRequests() {
	http.HandleFunc("/", HomePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("Hello World")

	handleRequests()
}
