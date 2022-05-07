package get

import (
	data "TaskManage/data"
	message "TaskManage/message"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//? get task route
func GetTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get task") //? link : 127.0.0.1:8080/getTask/1
	_params := mux.Vars(r)  //? get data from url like /1
	_flag := false
	for index, value := range data.TasksList {
		if _params["id"] == value.ID {
			json.NewEncoder(w).Encode(data.TasksList[index]) //? encode
			_flag = true
			break
		}
	}
	if _flag == false {
		message.ErrorMessage(w, r, "task not found")
	}
}
