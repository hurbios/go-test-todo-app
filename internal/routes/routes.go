package routes

import (
	"net/http"
	"test-todo-app/internal/handlers"

	"github.com/gorilla/mux"
)

func SetGetHeaders(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		f(w, r)
	}
}

func TodoRouter(todorouter *mux.Router) {
	todorouter.HandleFunc("/", SetGetHeaders(handlers.AllTodos)).Methods("GET")
	todorouter.HandleFunc("/{id}", SetGetHeaders(handlers.OneTodo)).Methods("GET")
}
