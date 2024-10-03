package main

import (
	"gojobquest/config"
	"gojobquest/router"
	"log"
	"net/http"
)

func main(){
	add:= ":9090"
	router:= router.NewRouter()
	config.DbConnect()


	log.Fatal(http.ListenAndServe(add,router))

}