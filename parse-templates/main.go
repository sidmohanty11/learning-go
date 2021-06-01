package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("text.html")

	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("new.html")
	//err = tpl.Execute(os.Stdout, nil) prints out the html
	err = tpl.Execute(nf, nil) // creates a new.html file -> copies tpl

	if err != nil {
		log.Fatalln(err)
	}
}
