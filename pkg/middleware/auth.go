package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func Login(store *sessions.CookieStore, f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username != "goder" || password != "alpha" {
			fmt.Println("login failed", username, password)
			f.ServeHTTP(w, r)
			return
		}
		fmt.Println("logged in")

		session.Values["authenticated"] = true
		session.Save(r, w)
		f.ServeHTTP(w, r)
	})
}

func Logout(store *sessions.CookieStore, f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")

		session.Values["authenticated"] = false
		session.Save(r, w)
		f.ServeHTTP(w, r)
	})
}

func Authenticate(store *sessions.CookieStore, f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		f.ServeHTTP(w, r)
	})
}
