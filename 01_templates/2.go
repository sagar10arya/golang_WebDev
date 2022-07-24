// templating with concatenation

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := `
	<!DOCTYPE HTML>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<body>
	<title>Hello World!</title>
	<h1> ` + name + ` </h1>
	</body>
	</head>`

	nf, err := os.Create("index.html") // create a new file(nf)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
