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
	// & -> &varName -> give me the address of anya
	anya.updateName("Malhotra")
	anya.details()
}

//* -> *pointer -> give me the value at that address
func (pointerToPerson *person) updateName(newLastName string) {
	//---------Rule to really remember-----------------:-
	//Turn address into value with *address / *pointer
	//Turn value into address with &value

	(*pointerToPerson).lastName = newLastName

	// *type -> *person -> looking for a pointer to a person(type description)
	// *pointer -> *pointerToPerson -> give me the value
}

func (p person) details() {
	fmt.Printf("%+v", p)
}
