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

var defaultMiddlewares = []string{"Logger", "Auth", "JsonResponse"}

var routes = Routes{
	Route{"Index", "GET", "/", Index, nil},
	// Route{"TodoIndex", "GET", "/api/todos", TodoIndex, defaultMiddlewares},
	// Route{"TodoShow", "GET", "/api/todos/{todo_id}", TodoShow, defaultMiddlewares},
	// Route{"TodoStore", "POST", "/api/todos", TodoStore, defaultMiddlewares},

	// Auth an invite
	Route{"Auth", "POST", "/api/auth", AuthHandler, []string{"Logger", "JsonResponse"}},

	// Update invite
	Route{"InviteUpdate", "PUT", "/api/invites/{invite_id}", InviteUpdate, defaultMiddlewares},
}
