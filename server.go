package main

import (
	"github.com/asaintgenis/todo/api"
	"github.com/asaintgenis/todo/db"
	"log"
	"net/http"
)

func main() {

	db.InitDB()
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
