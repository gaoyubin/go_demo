package main

import (
	"html/template"
	"log"
	"os"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {

	var str = `
			<html>
		<head>
		<title>Personal information</title>
		</head>
		<body style="text-align:center">
		<h3>Person general infor</h3>
		<hr>
		<ul>
		<li>Name: {{.Name}}:<p>
		<li>Id: {{.Id}} <p>
		<li>Country: {{.Country}}
		</ul>
		<hr>
		</body>
		</html>`
	p := Person{
		1,
		"low",
		12,
	}
	log.Println(p)
	tmpl := template.New("tmpl1")
	// tmpl.Parse("hello {{.Name}} welcome to")
	tmpl.Parse(str)

	tmpl.Execute(os.Stdout, p)
}
