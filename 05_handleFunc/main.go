package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Assemble Avengers")
}
func b(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Let's do it")
}

func main() {
	http.Handle("/avengers", http.HandlerFunc(a))
	http.Handle("/america/", http.HandlerFunc(b))

	http.ListenAndServe(":8080", nil)
}
