// Nested templates - Modularizing your code
/*
define:
	{{define "TemplateName"}}
	insert content here
	{{end}}
Use:
	{{template "TemplateName"}}
*/
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../nested/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
