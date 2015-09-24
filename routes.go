package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middlewares []string
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index, nil},
	Route{"Auth", "POST", "/auth", AuthHandler, []string{"Logger", "JsonResponse"}},
	Route{"TodoIndex", "GET", "/todos", TodoIndex, []string{"Logger", "Auth", "JsonResponse"}},
	Route{"TodoShow", "GET", "/todos/{todo_id}", TodoShow, []string{"Logger", "Auth", "JsonResponse"}},
	Route{"TodoStore", "POST", "/todos", TodoStore, []string{"Logger", "Auth", "JsonResponse"}},
}
