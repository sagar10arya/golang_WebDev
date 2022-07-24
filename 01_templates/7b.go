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

func init() {
	tpl = template.Must(template.ParseFiles("map.gohtml"))
}

// 2. Map

func main() {
	avengers := map[string]string{"Tony": "Iron-man", "Son of Odin": "Thor",
		"Soldier": "CaptainAmerica", "Witch": "Wanda", "Benner": "Hulk"}

	err := tpl.Execute(os.Stdout, avengers)
	if err != nil {
		log.Fatalln(err)
	}
}
