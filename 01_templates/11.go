// Predefined global Functions in templates

/*
Comparison:
eq: Returns the boolean truth of arg1 == arg2
ne: Returns the boolean truth of arg1 != arg2
lt: Returns the boolean truth of arg1 < arg2
le: Returns the boolean truth of arg1 <= arg2
gt: Returns the boolean truth of arg1 > arg2
ge: Returns the boolean truth of arg1 >= arg2
*/
package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("predefined_global_funcc.gohtml"))
}

func main() {
	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
