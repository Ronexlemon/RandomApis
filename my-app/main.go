package main

import (
	"fmt"
	"gomyapp/router"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	r := router.NewRouter()

	cors := handlers.CORS()(r)
	//start the http server
	http.Handle("/", r)
	fmt.Printf("starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", cors))

}
