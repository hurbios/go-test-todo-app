package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alltodos")
}

func TodoRouter(todorouter *mux.Router) {
	todorouter.HandleFunc("/", AllTodos)
}
