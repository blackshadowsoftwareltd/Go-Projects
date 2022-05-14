package apis

import (
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
