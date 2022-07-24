/*
When a form is submitted, we can pass the submitted values either through the request body payload or
through the URL.
If the form's method is post,then the values of the form are sent to the server through the request
body's payload.
If the form's method is get,the the values of the form are sent to thr server through the URL.
*/

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
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
	<input type="text" name="q">
	<input type="submit">
	</form><br>
	`+v)

}
