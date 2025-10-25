package main

import (
	"fmt"
	"net/http"
	"os"

	"test-todo-app/internal/routes"
	"test-todo-app/pkg/middleware"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	key   = []byte("store-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	fmt.Println("Good evening!")
	port := "8080"
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = envPort
	}

	fs := http.FileServer(http.Dir("web/static/"))

	r := mux.NewRouter()
	r.Use(middleware.SimpleLogger)
	r.Handle("/", fs)
	r.Handle("/login", middleware.Login(store, http.StripPrefix("/login", fs))).Methods("POST")
	r.Handle("/logout", middleware.Logout(store, http.StripPrefix("/logout", fs))).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	todorouter := r.PathPrefix("/api/todos").Subrouter()
	routes.TodoRouter(store, todorouter)

	fmt.Println("Server running on PORT", port)
	http.ListenAndServe(":"+port, r)
}
