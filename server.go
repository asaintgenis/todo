package main

import (
	"net/http"
	"github.com/asaintgenis/todo/api"
	"github.com/asaintgenis/todo/config"
	"log"
)

func main() {

	appConfig := config.AppConfig{PORT:"8080"}
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":"+appConfig.PORT, router))
}
