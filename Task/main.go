package main

import (
	"fmt"
	"gogroceries/router"
	"gomanagetask/configs"
	"log"
	"net/http"
)

//gotaskmanage

func main() {
	address := ":9090"
	router := router.NewRouter()
	configs.ConnectDB()

	fmt.Printf("listening on port:%s", address)
	log.Fatal(http.ListenAndServe(address, router))
}
