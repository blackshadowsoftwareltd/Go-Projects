package methods

import (
	databases "GoWithPostgre/database"
	dbFunc "GoWithPostgre/database/functions"
	"GoWithPostgre/messages"
	models "GoWithPostgre/models"
	"encoding/json"

	// messages "GoWithPostgre/messages"
	"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&body)

	// func PostMethod() {
	fmt.Println("Post Method")
	if databases.DB == nil {
		fmt.Println("Database is not connected")
		messages.ErrorMessage(w, r, "Failed to connect the Database.")
		return
	} else {
		fmt.Println("Database is connected")
	}
	///? write data to database
	err := dbFunc.AddUserToDB(body, w, r)
	if err != nil {
		return
	}
	///? read data from database
	var user models.UserModel
	user = dbFunc.GetLastUsersInfoFromDB()

	///? send data to client
	json.NewEncoder(w).Encode(user)

}
