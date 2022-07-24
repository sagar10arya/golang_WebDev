package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is my Cat Clara")
}
func b(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is my Dog Jimmy")
}

func main() {

	http.HandleFunc("/dog/", b)
	http.HandleFunc("/cat", a)

	http.ListenAndServe(":8080", nil)

}
