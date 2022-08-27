package methods

import (
	databases "GoWithPostgre/database"
	dbFunc "GoWithPostgre/database/functions"
	models "GoWithPostgre/models"

	// messages "GoWithPostgre/messages"
	"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

func PostMethod(w http.ResponseWriter, r *http.Request) {
	// func PostMethod() {
	fmt.Println("Post Method")
	if databases.DB == nil {
		fmt.Println("Database is not connected")
	} else {
		fmt.Println("Database is connected")
	}
	/*
			insertStatement := `insert into user_info values(3,'Name 0','me@io.io','pass 0','address 0','designation',0);`

			_, err := databases.DB.Exec(insertStatement)
			if err != nil {
				fmt.Println("Error", err)
				log.Fatal("Error inserting data", err)
				messages.ErrorMessage(w, r, "Error inserting data")
				return
			} else {
				fmt.Println("success")
				messages.Ok(w)
			}

		var res models.UserModel
		// var todos []string
		rows, err := databases.DB.Query("SELECT * FROM user_info")
		defer rows.Close()
		if err != nil {
			fmt.Println(err)
			log.Fatalln(err)
		}
		for rows.Next() {
			err = rows.Scan(&res.ID, &res.Name, &res.Email, &res.Password, &res.Address, &res.Designation, &res.Age)
			if err != nil {
				fmt.Println(err)
			}

			rows.Scan(&res)
			fmt.Println(res.Name)
			fmt.Println(res.ID)
			// 	todos = append(todos, res)
		}
		// fmt.Println("Todos : ", todos)
		// fmt.Println("Post Method End")
	*/

	var user models.UserModel
	user = dbFunc.GetLastUsersInfoFromDB()
	fmt.Println(user)

}
