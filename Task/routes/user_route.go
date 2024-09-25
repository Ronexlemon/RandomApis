package routes

import (
	"gomanagetask/controllers"

	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router){
	//All routes related to user comes here
	router.HandleFunc("/user",controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/user/{id}",controllers.GetAUser()).Methods("GET")
	router.HandleFunc("/user/{id}",controllers.Update()).Methods("PUT")
}