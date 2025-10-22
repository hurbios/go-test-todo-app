package main

import (
	"fmt"
	"net/http"

	"test-todo-app/internal/routes"
	"test-todo-app/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Good evening!")

	fs := http.FileServer(http.Dir("web/static/"))

	r := mux.NewRouter()
	r.Use(middleware.SimpleLogger)
	r.Handle("/", fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	todorouter := r.PathPrefix("/api/todos").Subrouter()
	routes.TodoRouter(todorouter)

	http.ListenAndServe(":8080", r)
}
