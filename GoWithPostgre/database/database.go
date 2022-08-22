package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "test"
)

var DB *sql.DB

var connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// var connStr = fmt.Sprintf("postgres://postgres:36166171@localhost/todos?sslmode=disable")

func CreateDatabase() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	DB = db
}
