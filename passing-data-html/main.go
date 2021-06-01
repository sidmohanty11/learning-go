package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	fName string
	lName string
	bDay  int
}

func (p person) happyBirthday() string {
	return "Happy Birthday to you!!!!"
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	p := person{
		fName: "Sidharth",
		lName: "Mohanty",
		bDay:  10032002,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", p)

	if err != nil {
		log.Fatalln(err)
	}
}
