package routes

import (
	"gojobquest/controller"

	"github.com/gorilla/mux"
)

func Quest(router *mux.Router){
	router.HandleFunc("/",controller.QuestCreate()).Methods("POST")

}