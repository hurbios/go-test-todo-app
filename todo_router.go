package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-todo-app/resources"

	"github.com/gorilla/mux"
)

func AllTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(resources.GetAllTodos())
}

func OneTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad id", http.StatusBadRequest)
		return
	}
	alltodos := resources.GetAllTodos()

	for i := 0; i < len(alltodos); i++ {
		if alltodos[i].Id == id {
			json.NewEncoder(w).Encode(alltodos[i])
			return
		}
	}
	http.Error(w, "Id not found", http.StatusNotFound)
}

func TodoRouter(todorouter *mux.Router) {
	todorouter.HandleFunc("/", AllTodos).Methods("GET")
	todorouter.HandleFunc("/{id}", OneTodo).Methods("GET")
}
