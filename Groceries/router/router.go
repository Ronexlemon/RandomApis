package router

import (
	"gogroceries/api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/gorceries",api.AllGroceries).Methods("GET")
	router.HandleFunc("/gorceries",api.GroceriesToBuy).Methods("POST")
	router.HandleFunc("/gorceries/{name}",api.SingleGroceries).Methods("GET")
	router.HandleFunc("/gorceries/{name}",api.UpdateGrocery).Methods("PUT")
	router.HandleFunc("/gorceries/{name}",api.DeleteGrocery).Methods("DELETE")


	return router
	
}