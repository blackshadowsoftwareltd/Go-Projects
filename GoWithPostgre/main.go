package main

import (
	databases "GoWithPostgre/database"
	router "GoWithPostgre/routes"
	"database/sql"
	"fmt"

	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {

	db, err := sql.Open("postgres", databases.ConnStr)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	DB = db
	router.HandleRoutes()
}
