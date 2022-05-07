package put

import (
	data "TaskManage/data"
	message "TaskManage/message"
	models "TaskManage/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//? update task route (2nd way)
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page") //? link : 127.0.0.1:8080/updateTask/1
	w.Header().Set("Content-Type", "application/json")
	_params := mux.Vars(r)                //? get data from url like /1
	_id, _ := strconv.Atoi(_params["id"]) //? id from paramiter
	var _task models.Task
	_ = json.NewDecoder(r.Body).Decode(&_task) //? decode (r.Body return body)
	_data := &data.TasksList[_id]

	if _task.TaskName != "" || _task.TaskDetail != "" {
		if _task.TaskName != "" {
			_data.TaskName = _task.TaskName
		}
		if _task.TaskDetail != "" {
			_data.TaskDetail = _task.TaskDetail
		}
		_data.Date = time.Now()
		data.TasksList[_id] = *_data
		message.SendMessage(w, r, "success")
	} else {
		message.SendMessage(w, r, "Task can't be empty")
	}
}

/*
//? update task route (1st way)
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page") //? link : 127.0.0.1:8080/updateTask/1
	w.Header().Set("Content-Type", "application/json")
	_params := mux.Vars(r) //? get data from url like /1
	var _task Tasks
	_ = json.NewDecoder(r.Body).Decode(&_task) //? decode (r.Body return body)
	_id, _ := strconv.Atoi(_params["id"])      //? id from paramiter
	_data := tasks[_id]

	if _task.TaskName != "" || _task.TaskDetail != "" {
		if _task.TaskName != "" {
			_data.TaskName = _task.TaskName
		}
		if _task.TaskDetail != "" {
			_data.TaskDetail = _task.TaskDetail
		}
		_data.Date = time.Now()
		tasks[_id] = _data
		sendMessage(w, r, "success")
	} else {
		sendMessage(w, r, "Task can't be empty")
	}
}
*/
