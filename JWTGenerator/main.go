package main

import (
	// jwtToken "RestJwt/token"
	handler "JWTGenerator/apis"
	"fmt"
)

func main() {
	fmt.Println("Started")

	handler.HandleRequests()
}
