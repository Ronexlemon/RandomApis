package routes

import (
	"gomanage/controller"

	"github.com/gorilla/mux"
)

func TaskRouter(router *mux.Router) {
	router.HandleFunc("/task", controller.TaskCreate()).Methods("POST")
	router.HandleFunc("/tasks/{id}", controller.TaskUserTasks()).Methods("GET")
	router.HandleFunc("/task/{id}", controller.TaskDelete()).Methods("DELETE")
	router.HandleFunc("/task/{id}", controller.TaskUpdate()).Methods("PUT")

}
