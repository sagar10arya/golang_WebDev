package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpeg">`)
}
func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpeg")
}
