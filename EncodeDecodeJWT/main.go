package main

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var UserId int = 1
var SecretKey string = "SecretKey"

func main() {
	///? Encode
	token, err := GenerateJWTToken(UserId)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println(token)

	///? Decode
	id, issuer, err := DecriptJWTTokenToId(token)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Printf("ID: %v, Issuer: %v", id, issuer)

}

///? Encode
func GenerateJWTToken(UserId int) (string, error) {
	fmt.Println("Generate Token")

	_id := strconv.Itoa(UserId)

	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256, //? method
		jwt.StandardClaims{
			Id:        _id,                                   //? ID
			Issuer:    _id,                                   //? Issuer
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //? Expire time
		},
	)

	token, err := claims.SignedString([]byte(SecretKey)) //? []byte(SecretKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil
}

///? Decode
func DecriptJWTTokenToId(tokenString string) (string, string, error) {
	fmt.Println("Decript Token")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil //? []byte(SecretKey)
	})

	if err != nil {
		fmt.Println("err ", err)
		return "", "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Id, claims.Issuer, nil
}
