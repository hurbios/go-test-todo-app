package main

import (
	"encoding/json"
	"net/http"
	"test-todo-app/resources"

	"github.com/gorilla/mux"
)

func AllTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(resources.GetAllTodos())
}

func TodoRouter(todorouter *mux.Router) {
	todorouter.HandleFunc("/", AllTodos).Methods("GET")
}
