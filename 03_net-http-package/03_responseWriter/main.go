package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Sagar-Arya", "Hey footballer!")
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintln(w, "<h1>Code in my func</h1>")

}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
