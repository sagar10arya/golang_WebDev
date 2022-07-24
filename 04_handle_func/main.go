package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the Iron-Man froomm Avengers")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the thor from Avengers")
}

func a(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the Captain America froomm Avengers")
}
func b(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is Scarlet Witch")
}

func main() {
	http.HandleFunc("/iron/", d)
	http.HandleFunc("/thor/", c)
	http.HandleFunc("/america/", a)
	http.HandleFunc("/witch/", b)

	http.ListenAndServe(":8080", nil)
}
