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
	tpl = template.Must(template.ParseFiles("slice-struct.gohtml"))
}

// 2. slice-Struct

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	gandhi := sage{
		Name:  "Gaandhi",
		Motto: "Be the change",
	}
	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}
	jesus := sage{
		Name:  "Jesus",
		Motto: "love all",
	}

	sages := []sage{buddha, gandhi, mlk, jesus}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
