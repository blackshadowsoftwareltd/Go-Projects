package get

import (
	data "TaskManage/data"
	"encoding/json"
	"fmt"
	"net/http"
)

//? get all tasks route
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All tasks") //? link : 127.0.0.1:8080/getAllTasks
	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(data.TasksList) //? encode
}
