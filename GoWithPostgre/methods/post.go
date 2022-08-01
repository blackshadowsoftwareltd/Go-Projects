package methods

import (
	databases "GoWithPostgre/database"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func PostMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Method")
	if databases.DB == nil {
		fmt.Println("Database is not connected")
	} else {
		fmt.Println("Database is connected")
	}
	_, err := databases.DB.Exec("INSERT INTO todos VALUES ($1)", "my")
	if err != nil {
		fmt.Println("Error", err)
		log.Fatal("Error inserting data", err)
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
	fmt.Println("Post Method End")
}
