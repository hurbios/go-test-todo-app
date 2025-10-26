package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func Login(store *sessions.CookieStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username != "goder" || password != "alpha" {
			fmt.Println("login failed", username, password)
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println("logged in")

		session.Values["authenticated"] = true
		session.Save(r, w)
		next.ServeHTTP(w, r)
	})
}

func Logout(store *sessions.CookieStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		session.Options.MaxAge = -1
		session.Values["authenticated"] = false
		session.Save(r, w)

		next.ServeHTTP(w, r)
	})
}

func Authenticate(store *sessions.CookieStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "session")

			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
