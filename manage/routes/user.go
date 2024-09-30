package routes

import (
	"gomanage/controller"

	"github.com/gorilla/mux"
)


func UserRouter(router *mux.Router){
	router.HandleFunc("/users", controller.Login()).Methods("GET")
	router.HandleFunc("/users", controller.SignUp()).Methods("POST")
	router.HandleFunc("/users/{id}", controller.Profile()).Methods("GET")
}

