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

type car struct {
	Manufacturer string
	Model        string
	Door         int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("struct-slice-struct.gohtml"))
}

// 2. Struct-slice-struct

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	f := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Door:         4,
	}
	c := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Door:         2,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
