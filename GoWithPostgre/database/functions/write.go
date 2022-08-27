package functions

import (
	databases "GoWithPostgre/database"
	"GoWithPostgre/messages"
	models "GoWithPostgre/models"
	"fmt"
	"log"
	"net/http"
)

func AddUserToDB(
	user models.UserModel,
	w http.ResponseWriter, r *http.Request) error {
	// insertStatement := `insert into user_info values(6,'Name 5','me@io.io','pass 0','address 0','designation',0);`

	_, err := databases.DB.Query("INSERT INTO user_info VALUES ($1, $2, $3, $4, $5, $6, $7)", 6, user.Name, user.Email, user.Password, user.Address, user.Designation, user.Age)
	if err != nil {
		fmt.Println("Error", err)
		log.Fatal("Error inserting data", err)
		messages.ErrorMessage(w, r, "Error inserting data")
		return err
	} else {
		fmt.Println("Data inserted successfully")
		return nil
	}
}
