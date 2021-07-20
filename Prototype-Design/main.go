package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// 1. Deep copying
// 2. Serialization

type Address struct {
	street, city, country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.street,
		a.city,
		a.country,
	}
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.address = p.address.DeepCopy()
	copy(q.friends, p.friends)
	return &q
}

func (p *Person) DeepCopySerialized() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	res := Person{}
	_ = d.Decode(&res)

	return &res
}

type Person struct {
	name    string
	address *Address
	friends []string
}

func main() {
	p := Person{"sid", &Address{"123", "Ctc", "IND"}, []string{"John", "Nani"}}

	// when you set john to sid, everything is copied from sid as IT IS A POINTER TO SID
	// so any changes made in john will affect sid as both of them share same pointer address.
	j := p
	j.name = "John"
	j.address.street = "321"

	fmt.Println(p.address, j.address) // address = 321 in both cases

	n := p
	n.address = &Address{
		p.address.street,
		p.address.city,
		p.address.country,
	}

	n.name = "Nani"
	n.address.street = "145"
	fmt.Println(p.address, n.address) // stays different as we made a completely new address

	// deep copying
	r := p.DeepCopy()
	r.name = "random"
	r.address.street = "999"
	fmt.Println(r, r.address)

	// serialization
	s := p.DeepCopy()
	s.name = "serialized"
	s.address.street = "100"
	s.friends = []string{"sid", "John"}
	fmt.Println(s, s.address)
}
