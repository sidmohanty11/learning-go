package main

import "fmt"

//Example-2
type person interface {
	talk() string
	eat() int
}

type male struct{}
type female struct{}

func main() {
	m := male{}
	f := female{}

	attributes(m)
	attributes(f)
}

func (m male) talk() string {
	return "Low Pitch"
}

func (f female) talk() string {
	return "High Pitch"
}

func (m male) eat() int {
	return 2400
}

func (f female) eat() int {
	return 1800
}

func attributes(p person) {
	fmt.Println(p.eat(), p.talk())
}
