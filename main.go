package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	AUTH_HEADER = "X-Auth"
)

func main() {

	port := os.Getenv("PORT")
	router := NewRouter()

	if port == "" {
		port = "8000"
	}

	fmt.Printf("Starting server on port %v\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
