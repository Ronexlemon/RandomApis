package main

import (
	"fmt"
	"gogroceries/router"
	"log"
	"net/http"
)

func main() {
	port := ":9090"
	r := router.NewRouter()

	r.Handle("/", r)
	fmt.Println("starting Server at port", port)

	log.Fatal(http.ListenAndServe(port, r))

}
