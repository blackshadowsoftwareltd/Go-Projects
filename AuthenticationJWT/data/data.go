package data

import (
	model "AuthenticationJWT/models"
)

var JwtKey = []byte("my_secret_key")
var Users []model.Credentials

func InitialiaeData() {
	user1 := model.Credentials{
		UserName: "user1",
		Password: "password1",
	}
	user2 := model.Credentials{
		UserName: "user2",
		Password: "password2",
	}
	user3 := model.Credentials{
		UserName: "user3",
		Password: "password3",
	}
	user4 := model.Credentials{
		UserName: "user4",
		Password: "password4",
	}

	Users = append(Users, user1)
	Users = append(Users, user2)
	Users = append(Users, user3)
	Users = append(Users, user4)
}
