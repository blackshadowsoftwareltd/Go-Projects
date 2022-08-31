package functions

import (
	databases "GoWithPostgre/database"
	"fmt"
)

///? create a table in database with a auto-increment primary key
func createDBTable(tableName string) error {
	///? CREATE TABLE table_name(); to create a table
	///? ALTER TABLE table_name ADD COLUMN id SERIAL PRIMARY KEY; to add a auto-increment primary key
	insertStatement := fmt.Sprintf("CREATE TABLE %s();ALTER TABLE %s ADD COLUMN id SERIAL PRIMARY KEY;", tableName, tableName)

	fmt.Println(insertStatement)
	_, err := databases.DB.Query(insertStatement)
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	return nil
}
