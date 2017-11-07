package main

import (
	"log"
	"net/http"

	"github.com/asaintgenis/todo/api"
	"github.com/asaintgenis/todo/db"
)

func main() {

	db.InitDB()
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
