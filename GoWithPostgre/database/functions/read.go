package functions

import (
	databases "GoWithPostgre/database"
	models "GoWithPostgre/models"
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
		//? password must be ignored. thats why, not use ampersand (&) before password field.
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Email, temp.Password, &temp.Address, &temp.Designation, &temp.Age)
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
		//? password must be ignored. thats why, not use ampersand (&) before password field.
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Email, temp.Password, &temp.Address, &temp.Designation, &temp.Age)
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
		//? password must be ignored. thats why, not use ampersand (&) before password field.
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Email, temp.Password, &temp.Address, &temp.Designation, &temp.Age)
		if err != nil {
			fmt.Println(err)
		}
		rows.Scan(&temp)
	}
	return temp
}
