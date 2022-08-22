package methods

import (
	databases "GoWithPostgre/database"
	// messages "GoWithPostgre/messages"
	"fmt"

	// "log"
	// "net/http"

	_ "github.com/lib/pq"
)

// func PostMethod(w http.ResponseWriter, r *http.Request) {
func PostMethod() {
	fmt.Println("Post Method")
	if databases.DB == nil {
		fmt.Println("Database is not connected")
	} else {
		fmt.Println("Database is connected")
	}
	insertStatement := `insert into test_table values(12,'Name 9');`
	_, err := databases.DB.Exec(insertStatement)
	if err != nil {
		fmt.Println("Error", err)
		// log.Fatal("Error inserting data", err)
		// messages.ErrorMessage(w, r, "Error inserting data")
		return
	} else {
		fmt.Println("success")
		// messages.Ok(w)
	}
	// var res string
	// var todos []string
	// rows, err := databases.DB.Query("SELECT * FROM todos")
	// defer rows.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	log.Fatalln(err)
	// }
	// for rows.Next() {
	// 	rows.Scan(&res)
	// 	todos = append(todos, res)
	// }
	// fmt.Println("Todos : ", todos)
	// fmt.Println("Post Method End")

}
