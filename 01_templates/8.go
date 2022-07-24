// function in templates
package main

import (
	"log"
	"os"
	"strings"
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

// Create a funcMap to register functions.
// "uc" is what the func will be called in template
// "uc" is the ToUpper func from package strings
// "ft" is the func I declaredtem
// "ft" slices a string, returning the first 3 characters

//Use funcMap to pass in your own functions to template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("funcTemplate.gohtml"))

}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

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
		Motto: "Hatred never ceases wtplith hatred but with love alone is healed",
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

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "funcTemplate.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
