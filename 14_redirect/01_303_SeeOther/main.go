package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your Request method at foo:", req.Method, "\n\n")
}
func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your Request method at bar:", req.Method, "\n\n")
	//process form submission here
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your Request method at barred:", req.Method, "\n\n")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
