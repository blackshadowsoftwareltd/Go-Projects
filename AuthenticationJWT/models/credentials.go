package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
