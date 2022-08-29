package functions

import (
	databases "GoWithPostgre/database"
	"fmt"
)

func createDBTable(tableName string) error {
	insertStatement := fmt.Sprintf("CREATE TABLE %s();ALTER TABLE %s ADD COLUMN id SERIAL PRIMARY KEY;", tableName, tableName)
	fmt.Println(insertStatement)
	_, err := databases.DB.Query(insertStatement)
	if err != nil {
		fmt.Println("Error", err)

		return err
	}
	return nil

}
