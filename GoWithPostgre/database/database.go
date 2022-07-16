package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = "8080"
	user     = "postgres"
	password = "36......"
	dbname   = "postgres"
)

var DB *sql.DB

var connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func CreateDatabase() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	DB = db
}
