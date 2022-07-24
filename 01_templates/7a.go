// passing composite data structures into templates

// 1. Slice (7a)
// 2. Map	(7b)
// 3. Struct	(7c)
// 4. slice-struct	(7d)
// 5. struct-slice-struct (7e)

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("slice.gohtml"))
}

// 1. Slice

func main() {
	avengers := []string{"Iron-man", "Thor", "CaptainAmerica", "Wanda", "Hulk"}

	err := tpl.Execute(os.Stdout, avengers)
	if err != nil {
		log.Fatalln(err)
	}
}
