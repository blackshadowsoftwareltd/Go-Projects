package token

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningkey = []byte("secret")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Println("Token creating")
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Remon"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningkey)

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return tokenString, nil

}
