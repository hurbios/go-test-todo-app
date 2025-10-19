package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Good evening!")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", fs)
	http.Handle("/night", fs)

	http.ListenAndServe(":80", nil)
}
