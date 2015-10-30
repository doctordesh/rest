package middlewares

import (
	"net/http"
)

func Auth(inner http.Handler, authHeader string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["X-Auth"] == nil {
			NewResponder(w).SendError(http.StatusBadRequest, "X-Auth header not present")
			return
		}

		inner.ServeHTTP(w, r)
	})
}
