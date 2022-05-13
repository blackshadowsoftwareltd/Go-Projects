package main

import (
	jwtToken "RestJwt/token"
	"fmt"
)

func main() {
	fmt.Println("Started")

	_token, _err := jwtToken.GenerateJWT()
	if _err != nil {
		fmt.Println("Error:", _err)
	} else {
		fmt.Println("Created Token:", _token)
	}
}
