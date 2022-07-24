// templating with concatenation
package main

import "fmt"

func main() {
	name := "Sagar Arya"
	tpl := `
	<!DOCTYPE HTML>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<body>
	<title>Hello World!</title>
	<h1> ` + name + ` </h1>
	</body>
	</head>`

	fmt.Println(tpl)
}
