package main

import (
	// jwtToken "RestJwt/token"
	handler "RestJwt/apis"
	"fmt"
)

func main() {
	fmt.Println("Started")

	handler.HandleRequests()
}
