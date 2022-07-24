// passing composite data structures into templates

// 1. Slice
// 2. Map
// 3. Struct
// 4. slice-struct
// 5. struct-slice-struct

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("struct.gohtml"))
}

// 3. Struct

func main() {
	buddha := sage{Name: "Buddha", Motto: "The belief of no beliefs"}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
