package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		for _, middleware := range route.Middlewares {
			if middleware == "Logger" {
				handler = Logger(handler, route.Name)
			}
			if middleware == "JsonResponse" {
				handler = JsonResponse(handler)
			}
			if middleware == "Auth" {
				handler = Auth(handler, AUTH_HEADER)
			}
		}

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}

func sendError(code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"status": http.StatusText(code),
		"code":   strconv.Itoa(code),
	})
}
