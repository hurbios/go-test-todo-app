package main

import (
	"fmt"
	"net/http"
	"os"

	"test-todo-app/internal/routes"
	"test-todo-app/pkg/middleware"
	"test-todo-app/pkg/middleware/auth"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type App struct {
	Router *mux.Router
	Store  *sessions.CookieStore
}

func main() {
	fmt.Println("Good evening!")

	app := &App{
		Router: mux.NewRouter(),
		Store:  sessions.NewCookieStore([]byte("a-secret-session-key")),
	}

	port := "8080"
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = envPort
	}

	fs := http.FileServer(http.Dir("web/static/"))

	app.Router.Use(middleware.SimpleLogger)
	app.Router.Handle("/", fs)
	app.Router.Handle("/login", auth.Login(app.Store, http.StripPrefix("/login", fs))).Methods("POST")
	app.Router.Handle("/logout", auth.Logout(app.Store, http.StripPrefix("/logout", fs))).Methods("POST")
	app.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	todorouter := app.Router.PathPrefix("/api/todos").Subrouter()
	todorouter.Use(auth.Authenticate(app.Store))
	routes.TodoRouter(todorouter)

	fmt.Println("Server running on PORT", port)
	http.ListenAndServe(":"+port, app.Router)
}
