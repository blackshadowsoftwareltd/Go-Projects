package main

import (
	data "TaskManage/data"
	handle "TaskManage/handler"

	"fmt"
)

func main() {

	fmt.Println("started")
	data.AllTaskList()
	//?

	handle.HandlRoutes()
}
