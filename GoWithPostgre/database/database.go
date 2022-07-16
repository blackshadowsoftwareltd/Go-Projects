package database

import "fmt"

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// )

const (
	host     = "localhost"
	port     = "8080"
	user     = "postgres"
	password = "36......"
	dbname   = "postgres"
)

var ConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// onnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// func ConnectToDatabase() (d  *sql.DB ) {
// 	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)

// 	}

// }
