package api

import (
	"net/http"
)

//Route route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes routes
type Routes []Route

var routes = Routes{
	Route{
		"GetTodos",
		"GET",
		"/todos",
		GetTodos,
	},
	Route{
		"GetTodo",
		"GET",
		"/todos/{todoId}",
		GetTodo,
	},
}
