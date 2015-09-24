package main

import (
	"encoding/json"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := authUser(r.PostForm["email"][0], r.PostForm["password"][0])

	if err != nil {
		sendError(http.StatusUnauthorized, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}
