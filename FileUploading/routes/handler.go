package routes

import (
	methods "FileUploading/uploads"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter()
	///? routes
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/upload", methods.FileUpload).Methods("POST")

	///?
	log.Fatal(http.ListenAndServe(":8080", router))
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Endpoint Hit")
}
