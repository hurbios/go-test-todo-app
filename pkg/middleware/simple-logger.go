package middleware

import (
	"fmt"
	"net/http"
)

func SimpleLogger(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, "Request received with path ", r.URL.Path)
		f.ServeHTTP(w, r)
	})
}
