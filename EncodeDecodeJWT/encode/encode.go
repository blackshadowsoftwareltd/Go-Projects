package encode

import (
	"fmt"
	"strconv"
	"time"
data "EncodeDecodeJWT/data"
	jwt "github.com/dgrijalva/jwt-go"
)

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

	token, err := claims.SignedString([]byte(data.SecretKey)) //? []byte(SecretKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil
}
