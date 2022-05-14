package apis

import (
	jwtToken "JWTGenerator/token"
	"fmt"
	"net/http"
	// "time"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	_token, _err := jwtToken.GenerateJWT()
	if _err != nil {
		fmt.Println("Error:", _err)
	} else {
		fmt.Fprintln(w, _token)                 //? Fprintln is used to print to the web page
		fmt.Println("Token created : ", _token) //? Println is used to print to the console
	}

}
