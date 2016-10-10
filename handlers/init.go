package handlers

import (
	"fmt"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("asoijd")
}

func init() {
	Handler(a, b)
}
