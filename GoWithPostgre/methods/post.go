package methods

import (
	"fmt"
	"net/http"
)

func PostMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Method")
}
