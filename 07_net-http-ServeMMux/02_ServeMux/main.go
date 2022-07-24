/*
In go ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request
against a list of registered patterns and calls the handler for the pattern that most closely
matches the URL.*/
package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is Johny Dogs")
}

func (d hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is Tommy Tommy")
}

func main() {
	var a hotdog
	var b hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", a)
	mux.Handle("/cat", b)

	http.ListenAndServe(":8080", mux)
}
