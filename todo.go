package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "ASD"},
		Todo{Name: "ASD"},
	}

	if err := sendJson(w, todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todo_id"]
	fmt.Fprintf(w, "<h1>Todo#show %q</h1>", todoId)
}

func TodoStore(w http.ResponseWriter, r *http.Request) {
	todo := Todo{Name: "aoijsd"}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(todo)
}
