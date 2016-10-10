package handlers

import (
	"fmt"
	"net/http"
)

func AuthIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lorem ipsum")
}
