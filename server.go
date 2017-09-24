package main

import (
	"net/http"
	"github.com/asaintgenis/todo/api"
	"log"
)

func main() {

	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
