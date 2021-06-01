package main

import "fmt"

//to whom it may concern -> our prog has a new type called bot
//interface -> abstract outlook
//concrete types -> map,int,string,struct,englishBot etc
//interface type -> bot
//can't declare a value directly to interface

//*INTERFACES ARE A CONTRACT TO HELP US MANAGE TYPES!*

//Example-1
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "hello, there!"
}

func (spanishBot) getGreeting() string {
	return "hola amigos!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
