package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your Request method at foo:", req.Method, "\n\n")
}
func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your Request method at bar:", req.Method, "\n\n")
	//process form submission here
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}
