package main

import (
	// "fmt"
	"log"
	"net/http"
	// "time"

	// jwt "github.com/dgrijalva/jwt-go"
	handler "AuthenticationJWT/handler"
	data "AuthenticationJWT/data"
)

func main() {
	Handleer()
}



func Handleer() {
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/home", handler.Home)
	http.HandleFunc("/refresh", handler.Refresh)
	log.Fatal(http.ListenAndServe(":8080", nil))

	data.InitialiaeData()
}
