package delete

import (
	data "TaskManage/data"
	message "TaskManage/message"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//? delete task route
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page") //? link : 127.0.0.1:8080/deleteTask/1
	w.Header().Set("Content-Type", "application/json")
	_id := mux.Vars(r)["id"] //? get data from url like /1
	_flag := false
	for index, value := range data.TasksList {
		if value.ID == _id {
			data.TasksList = append(data.TasksList[:index], data.TasksList[index+1:]...)
			_flag = true
			message.SendMessage(w, r, "success")
			break
		}
	}
	if _flag == false {
		message.ErrorMessage(w, r, "Invalid id")
	}
}
