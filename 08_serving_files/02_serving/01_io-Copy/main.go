package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dogs.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/dogs.jpeg">`)

}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("dogs.jpeg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
