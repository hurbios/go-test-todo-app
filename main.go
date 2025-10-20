package main

import (
	"fmt"
	"net/http"

	"test-todo-app/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Good evening!")

	fs := http.FileServer(http.Dir("static/"))

	r := mux.NewRouter()
	r.Use(middlewares.SimpleLogger)
	r.Handle("/", fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	todorouter := r.PathPrefix("/api/todos").Subrouter()
	TodoRouter(todorouter)

	http.ListenAndServe(":80", r)
}
