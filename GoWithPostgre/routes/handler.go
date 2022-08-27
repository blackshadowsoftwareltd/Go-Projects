package routes

import (
	methods "GoWithPostgre/methods"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter()
	///? routes
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/post", methods.PostMethod).Methods("POST")
	// methods.PostMethod()

	///?
	port := os.Getenv("PORT")
	if port == "" {
		port = "5431"
	}
	port = "5431"
	log.Fatal(http.ListenAndServe(":"+port, router))
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Endpoint Hit")
}
