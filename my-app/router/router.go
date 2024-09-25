package router

import (
	"gomyapp/api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r:= mux.NewRouter()

	r.HandleFunc("/words/{search}",api.GetWord).Methods("GET")
	return r

}