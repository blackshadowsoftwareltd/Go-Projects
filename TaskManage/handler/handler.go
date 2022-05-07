package handler

import (
	delete "TaskManage/methods/delete"
	get "TaskManage/methods/get"
	post "TaskManage/methods/post"
	put "TaskManage/methods/put"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//? create routes
func HandlRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", get.HomePage).Methods("GET")                        //? link : 127.0.0.1:8080/
	router.HandleFunc("/getTask/{id}", get.GetTask).Methods("GET")             //? link : 127.0.0.1:8080/getTask/1
	router.HandleFunc("/getAllTasks", get.GetAllTasks).Methods("GET")          //? link : 127.0.0.1:8080/getAllTasks
	router.HandleFunc("/createTask", post.CreateTask).Methods("POST")          //? link : 127.0.0.1:8080/createTask
	router.HandleFunc("/updateTask/{id}", put.UpdateTask).Methods("PUT")       //? link : 127.0.0.1:8080/updateTask/1
	router.HandleFunc("/deleteTask/{id}", delete.DeleteTask).Methods("DELETE") //? link : 127.0.0.1:8080/deleteTask/1
	//?
	log.Fatal(http.ListenAndServe(":8080", router))
}
