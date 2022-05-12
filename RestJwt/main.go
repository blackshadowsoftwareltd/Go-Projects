package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {}

func GenerateJWT() (string error) {
	token := jwt.New(jwt.SigningMethodES256)
	fmt.Println("token:", token)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Remon"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

}
