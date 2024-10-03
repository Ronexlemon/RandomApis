package main

import (
	"gojobquest/router"
	"log"
	"net/http"
)

func main(){
	add:= ":9090"
	router:= router.NewRouter()


	log.Fatal(http.ListenAndServe(add,router))

}