package main

import (
	databases "GoWithPostgre/database"
	router "GoWithPostgre/routes"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Started")
	databases.CreateDatabase()
	router.HandleRoutes()
}
