package main

import (
	"fmt"
	"gogroceries/router"
	"log"
	"net/http"
)

//gotaskmanage

func main() {
	address := ":9090"
	router := router.NewRouter()

	fmt.Printf("listening on port:%s", address)
	log.Fatal(http.ListenAndServe(address, router))
}
