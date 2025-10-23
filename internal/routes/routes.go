package routes

import (
	"net/http"
	"test-todo-app/internal/handlers"
	"test-todo-app/pkg/middleware"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func SetGetHeaders(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		f(w, r)
	}
}

func TodoRouter(store *sessions.CookieStore, todorouter *mux.Router) {
	todorouter.HandleFunc("/", middleware.Authenticate(store, SetGetHeaders(handlers.AllTodos))).Methods("GET")
	todorouter.HandleFunc("/{id}", middleware.Authenticate(store, SetGetHeaders(handlers.OneTodo))).Methods("GET")
}
