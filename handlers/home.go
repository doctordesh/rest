package handlers

import (
	"fmt"
	"net/http"
)

func Something() http.HandlerFunc {
	return Index
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
