package message

import (
	"encoding/json"
	"net/http"
)

///? send message
func SendMessage(w http.ResponseWriter, r *http.Request, message string) {
	json.NewEncoder(w).Encode(map[string]string{"message": message}) //? success message
}

//? error message
func ErrorMessage(w http.ResponseWriter, r *http.Request, message string) {
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
