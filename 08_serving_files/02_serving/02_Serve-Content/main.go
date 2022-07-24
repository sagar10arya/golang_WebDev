package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset:utf-8")
	io.WriteString(w, `<img src="toby.jpeg">`)
}
func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}

// func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
