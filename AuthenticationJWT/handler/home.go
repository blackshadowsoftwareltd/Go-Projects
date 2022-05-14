package handler

import (
	"AuthenticationJWT/data"
	model "AuthenticationJWT/models"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_cookie, _err := r.Cookie("token")
	if _err != nil {
		if _err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_tokenString := _cookie.Value
	_claims := &model.Claims{}
	_token, _xErr := jwt.ParseWithClaims(_tokenString, _claims, func(t *jwt.Token) (interface{}, error) {
		return data.JwtKey, nil
	})
	if _xErr != nil {
		if _xErr == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !_token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
	}
	_message := fmt.Sprint("Welcome ", _claims.UserName)
	w.Write([]byte(_message))
}
