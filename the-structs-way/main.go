package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	anya := person{
		firstName: "Anya",
		lastName:  "Chalotra",
		contact: contactInfo{
			email: "anyachalotra@gmail.com",
			zip:   68940,
		},
	}
	anya.updateName("Malhotra")
	anya.details()
}

func (p person) updateName(newLastName string) {
	p.lastName = newLastName
}

func (p person) details() {
	fmt.Printf("%+v", p)
}
