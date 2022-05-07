package models

import (
	"time"
)

//? struct
type Task struct {
	ID         string    `json:"id"`
	TaskName   string    `json:"taskName"`
	TaskDetail string    `json:"taskDetail"`
	Date       time.Time `json:"date"`
}
