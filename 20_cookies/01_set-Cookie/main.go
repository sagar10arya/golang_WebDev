package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "Can I tell you a Secret?",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER BUDDY")
	fmt.Fprintln(w, "in chrome go to: devtools/application/cookies")
}
func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "Your Cookie: ", c)
}
