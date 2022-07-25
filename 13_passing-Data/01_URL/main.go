package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Do my Search: "+v)
}
