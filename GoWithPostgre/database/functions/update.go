package functions

import (
	databases "GoWithPostgre/database"
	models "GoWithPostgre/models"
	"fmt"
	"net/http"
	"strings"
)

func updateDBTable(user models.UserModel) error {
	return nil
}
func UpdateUserInfoDBTable(tableName string, user models.UserModel, w http.ResponseWriter, r *http.Request) error {
	query := fmt.Sprintf("UPDATE %s SET ", tableName) ///? first insert the table name
	if user.Name != "" {
		query += fmt.Sprintf("name = '%s',", user.Name) ///? then insert the name
	}
	if user.Email != "" {
		query += fmt.Sprintf(" email = '%s',", user.Email) ///? then insert the email
	}
	if user.Address != "" {
		query += fmt.Sprintf(" address = '%s',", user.Address) ///? then insert the address
	}
	if user.Designation != "" {
		query += fmt.Sprintf(" designation = '%s',", user.Designation) ///? then insert the designation
	}
	if user.Age != 0 {
		query += fmt.Sprintf(" age = %d ", user.Age) ///? then insert the age
	}
	query = fmt.Sprintf(strings.TrimSuffix(query, ",")) ///? It will remove the last comma

	query += fmt.Sprintf(" WHERE id = %d;", user.Id) ///? last check the condition where id == user.Id
	fmt.Println(query)

	_, err := databases.DB.Query(query)
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	return nil
}

// ?, email=?, password=?, address=?, designation=?, age=?
