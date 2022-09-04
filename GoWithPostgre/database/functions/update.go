package functions

import (
	models "GoWithPostgre/models"
	"fmt"
	"net/http"
)

func updateDBTable(user models.UserModel) error {
	return nil
}
func UpdateUserInfoDBTable(tableName string, user models.UserModel, w http.ResponseWriter, r *http.Request) error {
	insertStatement := fmt.Sprintf("UPDATE %s SET name= xyz...	WHERE id=%d;", tableName, user.Id)
	fmt.Println(insertStatement)
	fmt.Println(user)
	// _, err := databases.DB.Query(insertStatement)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return err
	// }
	return nil
}

// ?, email=?, password=?, address=?, designation=?, age=?
