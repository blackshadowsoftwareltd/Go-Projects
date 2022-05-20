package decode

import (
	data "EncodeDecodeJWT/data"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

///? Decode
func DecriptJWTTokenToId(tokenString string) (string, string, error) {
	fmt.Println("Decript Token")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(data.SecretKey), nil //? []byte(SecretKey)
	})

	if err != nil {
		fmt.Println("err ", err)
		return "", "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Id, claims.Issuer, nil
}
