package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/asaintgenis/todo/model"
	"github.com/asaintgenis/todo/service"
	"github.com/gorilla/mux"
)

// GetTodos return a JSON of all todos or an error code
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

//GetTodo return the selected todo in JSON or an error message
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

// PostTodo create the todo send in body
func PostTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo model.Todo
	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	service.PostTodo(&todo)
}

// PutTodo update the todo send in body
func PutTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo model.Todo
	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	service.PutTodo(&todo)
}

// DeleteTodo delete the todo or send an error message
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		log.Fatal("ID must be an unsigned integer")
	}
	err = service.DeleteTodo(uint(todoID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w)
}
