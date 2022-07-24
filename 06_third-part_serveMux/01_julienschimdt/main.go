package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.new()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/category/:article", blogRead)
	mux.POST("/blog/category/:article", blogWrite)

	http.ListenAndServe(":8080", mux)

}

func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "USER, %s!\n", ps.Byname("name"))
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "Read Category, %s!\n", ps.Byname("category"))
	fmt.Fprintln(w, "Read Article, %s!\n", ps.Byname("article"))
}

func blogWrite(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "Write Category, %s!\n", ps.Byname("category"))
	fmt.Fprintln(w, "Write Article, %s!\n", ps.Byname("article"))
}
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := template.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}
func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := template.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}
func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := template.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}
func apply(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := template.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}
func applyProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := template.ExecuteTemplate(w, "applyProcess	.gohtml", nil)
	HandleError(w, err)
}
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
