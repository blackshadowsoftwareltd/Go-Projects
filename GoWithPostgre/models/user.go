package models

type UserModel struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	Designation string `json:"designation"`
	Age         int    `json:"age"`
}

type UserModelResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Designation string `json:"designation"`
	Age         int    `json:"age"`
}
