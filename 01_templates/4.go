// package text/template : parsing and executing templates
/*
=>template.Must(t *template, err error)-> takes template pointer & error as input and
it's check the error and return back the template pointer.
=>.ParseGlob("templates/*")  , It's mean pass all the file in the template dir.
  .ParseGlob("templates/*.gohtml") , It's mean pass all the file in the template
   dir which have extension of .gohtml.
=>.ExecuteTemplat(writer,name.data) - it allow to execute specific container.
=>
*/
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*"))
}

func main() {
	// tpl, err := template.ParseGlob("templates/*")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
