package handler

import (
	data "AuthenticationJWT/data"
	model "AuthenticationJWT/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: login")
	var credentials model.Credentials
	_isExist := false
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, _user := range data.Users {
		fmt.Printf("%v -- %v -- %v -- %v", _user.UserName, credentials.UserName, _user.Password, credentials.Password)

		if _user.UserName == credentials.UserName && _user.Password == credentials.Password {
			fmt.Println("User found")
			_isExist = true
			break
		}
	}
	fmt.Println(_isExist)
	if _isExist == false {
		fmt.Println("Unauthorized access")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println(4)
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &model.Claims{
		UserName: credentials.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println(5)
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_tokenString, err := _token.SignedString([]byte(data.JwtKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(6)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   _tokenString,
		Expires: expirationTime,
	})

}
