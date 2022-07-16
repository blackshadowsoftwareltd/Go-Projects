package routes

import (
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

	///?
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Endpoint Hit")
}
