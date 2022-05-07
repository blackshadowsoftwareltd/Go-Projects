package post

import (
	data "TaskManage/data"
	message "TaskManage/message"
	models "TaskManage/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//? create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
	w.Header().Set("Content-Type", "application/json")
	var _task models.Task
	_ = json.NewDecoder(r.Body).Decode(&_task) //? decode (r.Body return body)
	_taskLength := len(data.TasksList)         //? task length
	_task.ID = strconv.Itoa(_taskLength)       //? initialize a incremented id
	_task.Date = time.Now()
	if _task.TaskName == "" {
		message.ErrorMessage(w, r, "Task name can't be empty")
	} else if _task.TaskDetail == "" {
		message.ErrorMessage(w, r, "Task detail can't be empty")
	} else {
		data.TasksList = append(data.TasksList, _task) //? add new task

		if len(data.TasksList) != _taskLength { //? that means, new task add in the slice
			message.SendMessage(w, r, "success")
		} else {
			message.SendMessage(w, r, "failed")
		}
	}
}
