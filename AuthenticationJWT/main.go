package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"

	// "time"

	// jwt "github.com/dgrijalva/jwt-go"
	data "AuthenticationJWT/data"
	handler "AuthenticationJWT/handler"
)

func main() {
	Handleer()
}



func Handleer() {
	fmt.Println("Started")
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/home", handler.Home)
	http.HandleFunc("/refresh", handler.Refresh)
	log.Fatal(http.ListenAndServe(":8080", nil))

	data.InitialiaeData()
}
