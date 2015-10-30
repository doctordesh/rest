package main

import (
	"net/http"

	"github.com/doctordesh/rest/middlewares"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		for _, middleware := range route.Middlewares {
			if middleware == "Logger" {
				handler = middlewares.Logger(handler, route.Name)
			}
			if middleware == "JsonResponse" {
				handler = middlewares.JsonResponse(handler)
			}
			if middleware == "Auth" {
				handler = middlewares.Auth(handler, AUTH_HEADER)
			}
		}

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}
