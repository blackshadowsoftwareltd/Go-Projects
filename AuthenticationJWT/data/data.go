package data

import (
	model "AuthenticationJWT/models"
)

var JwtKey = []byte("my_secret_key")
var Users []model.Credentials

func InitialiaeData() {
	user1 := model.Credentials{
		UserName: "user 1",
		Password: "pasword 1",
	}
	user2 := model.Credentials{
		UserName: "user 2",
		Password: "password 2",
	}
	user3 := model.Credentials{
		UserName: "user 3",
		Password: "password 3",
	}
	user4 := model.Credentials{
		UserName: "user 4",
		Password: "password 4",
	}

	Users = append(Users, user1)
	Users = append(Users, user2)
	Users = append(Users, user3)
	Users = append(Users, user4)
}
