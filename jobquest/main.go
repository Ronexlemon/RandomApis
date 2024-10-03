package main

import (
	"fmt"
	"gojobquest/config"
	"gojobquest/router"
	"gojobquest/routes"
	"log"
	"net/http"
)

func main(){
	add:= ":9090"
	router:= router.NewRouter()
	config.DbConnect()
	routes.Quest(router)

fmt.Println("Connected at port",add)
	log.Fatal(http.ListenAndServe(add,router))

}