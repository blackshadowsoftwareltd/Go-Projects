package handler

import (
	data "AuthenticationJWT/data"
	model "AuthenticationJWT/models"
	"encoding/json"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials
	_isExist := false
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, _user := range data.Users {
		if _user.UserName == credentials.UserName && _user.Password == credentials.Password {
			_isExist = true
			break
		}
	}
	if _isExist == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &model.Claims{
		UserName: credentials.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_tokenString, err := _token.SignedString([]byte(data.JwtKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   _tokenString,
		Expires: expirationTime,
	})

}
