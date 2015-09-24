package main

import (
	"net/http"
)

func Auth(inner http.Handler, authHeader string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["X-Auth"] == nil {
			sendError(http.StatusBadRequest, w)
			return
		}

		inner.ServeHTTP(w, r)
	})
}