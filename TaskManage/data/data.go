package data

import (
	models "TaskManage/models"
	"fmt"
	"time"
)

var TasksList []models.Task

//? initialize tasks
func AllTaskList() {
	task1 := models.Task{
		ID:         "0",
		TaskName:   "Task 1",
		TaskDetail: "Task 1 Detail",
		Date:       time.Now(),
	}
	task2 := models.Task{
		ID:         "1",
		TaskName:   "Task 2",
		TaskDetail: "Task 2 Detail",
		Date:       time.Now(),
	}

	TasksList = append(TasksList, task1)
	TasksList = append(TasksList, task2)
	fmt.Println(TasksList)
}
