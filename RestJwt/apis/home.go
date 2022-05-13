package apis

import (
	jwtToken "RestJwt/token"
	"fmt"
	"net/http"
	// "time"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	_token, _err := jwtToken.GenerateJWT()
	if _err != nil {
		fmt.Println("Error:", _err)
	} else {
		fmt.Println("Token :", _token)
	}

}
