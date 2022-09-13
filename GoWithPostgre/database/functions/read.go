package functions

import (
	databases "GoWithPostgre/database"
	models "GoWithPostgre/models"
	"database/sql"
	"fmt"
	"log"
)

///? Get All Users info from database
func GetAllUsersFromDB() []models.UserModel {
	rows, err := databases.DB.Query("SELECT * FROM user_info")
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	defer rows.Close()
	var list []models.UserModel
	var temp models.UserModel
	for rows.Next() {
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Email, &temp.Password, &temp.Address, &temp.Designation, &temp.Age)
		if err != nil {
			fmt.Println(err)
		}
		rows.Scan(&temp)
		list = append(list, temp)
	}
	return list
}

///? Get only Single column info from database
func GetSingleUserInfoFromDB(id int64) models.UserModel {
	rows, err := databases.DB.Query("SELECT * FROM user_info WHERE ID = $1", id)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	defer rows.Close()
	var temp models.UserModel
	for rows.Next() {
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Email, &temp.Password, &temp.Address, &temp.Designation, &temp.Age)
		if err != nil {
			fmt.Println(err)
		}
		rows.Scan(&temp)
	}
	return temp
}

///? Get All Users info from database
func GetLastUsersInfoFromDB() models.UserModel {
	rows, err := databases.DB.Query("SELECT * FROM user_info ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	defer rows.Close()
	var temp models.UserModel
	for rows.Next() {
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Email, &temp.Password, &temp.Address, &temp.Designation, &temp.Age)
		if err != nil {
			fmt.Println(err)
		}
		rows.Scan(&temp)
	}
	return temp
}
func UserIsExistInDB(id int64, email string) int64 { ///? This method can check whether the user is exist or not by id and email.
	var rows *sql.Rows
	var err error
	if id != 0 {
		rows, err = databases.DB.Query("SELECT * FROM user_info WHERE ID = $1", id)
	} else if email != "" {
		rows, err = databases.DB.Query("SELECT * FROM user_info WHERE email = $1", email)
	}

	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	defer rows.Close()
	var temp models.UserModel
	for rows.Next() {
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Email, &temp.Password, &temp.Address, &temp.Designation, &temp.Age)
		if err != nil {
			fmt.Println(err)
		}
		rows.Scan(&temp)
	}
	fmt.Println(temp.Id)
	if id != 0 && temp.Id == id {
		return temp.Id
	}
	if email != "" && temp.Email == email {
		return temp.Id
	}
	return 0

}
