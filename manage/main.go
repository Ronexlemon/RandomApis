package main

import (
	"fmt"
	"gomanage/config"
	"gomanage/router"
	"gomanage/routes"
	"log"
	"net/http"
)

func main(){
	port := ":9090"
	//router
	router := router.NewRouter()
	routes.UserRouter(router)
	routes.TaskRouter(router)
	config.NewClient()
	fmt.Println("Listening to port",port)
	log.Fatal(http.ListenAndServe(port,router))

}