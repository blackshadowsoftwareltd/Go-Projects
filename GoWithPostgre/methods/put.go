package methods

import (
	databases "GoWithPostgre/database"
	dbFunc "GoWithPostgre/database/functions"
	"GoWithPostgre/messages"
	models "GoWithPostgre/models"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Method")
	var body models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&body)

	fmt.Println("Post Method")
	if databases.DB == nil {
		fmt.Println("Database is not connected")
		messages.ErrorMessage(w, r, "Failed to connect the Database.")
		return
	} else {
		fmt.Println("Database is connected")
	}
	///? write data to database
	err := dbFunc.UpdateUserInfoDBTable("user_info", body, w, r)
	if err != nil {
		return
	}
	///? read data from database
	user := dbFunc.GetLastUsersInfoFromDB()

	///? send data to client
	responseBody := models.UserModelResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Address:     user.Address,
		Designation: user.Designation,
		Age:         user.Age,
	}
	json.NewEncoder(w).Encode(responseBody)

}
