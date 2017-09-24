package api

import (
	"encoding/json"
	"fmt"
	"github.com/asaintgenis/todo/model"
	"github.com/asaintgenis/todo/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

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

func PostTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo model.Todo
	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	service.PostTodo(&todo)
}

func PutTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo model.Todo
	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	service.PutTodo(&todo)
}

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
