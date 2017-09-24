package api

import (
	"encoding/json"
	"fmt"
	"github.com/asaintgenis/todo/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//CatsIndex send a json back with all the cats database
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := service.GetTodos()
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todos); err != nil {
			panic(err)
		}
		fmt.Fprintln(w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
	fmt.Fprintln(w)
}

//CatShow show a specific required beautiful cat
func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		log.Fatal("ID must be an unsigned integer")
	}
	todo, err := service.GetTodo(uint(todoID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		fmt.Fprintln(w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
	fmt.Fprintln(w)
}
