package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(w, "Any code you want in this function")
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Hey doggi !")

	case "/cat":
		io.WriteString(w, "Hey kitty !")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
