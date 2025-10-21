package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-todo-app/resources"

	"github.com/gorilla/mux"
)

func SetGetHeaders(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		f(w, r)
	}
}

func AllTodos(w http.ResponseWriter, r *http.Request) {
	cat := r.URL.Query().Get("cat")
	alltodos := resources.GetAllTodos()

	if cat != "" {
		var alltodosincategory []resources.Todo
		for i := 0; i < len(alltodos); i++ {
			if alltodos[i].Category == cat {
				alltodosincategory = append(alltodosincategory, alltodos[i])
			}
		}
		json.NewEncoder(w).Encode(alltodosincategory)
	} else {
		json.NewEncoder(w).Encode(alltodos)
	}
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
	todorouter.HandleFunc("/", SetGetHeaders(AllTodos)).Methods("GET")
	todorouter.HandleFunc("/{id}", SetGetHeaders(OneTodo)).Methods("GET")
}
