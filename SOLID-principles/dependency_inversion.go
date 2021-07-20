package main

import "fmt"

// DEPENDENCY INVERSION PRINCIPLE

// High level modules should not depend on Low level modules
// Both should depend on abstractions ...

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// and other stuffs
}

type Info struct {
	from *Person
	rel  Relationship
	to   *Person
}

// low level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	res := make([]*Person, 0)

	for i, v := range r.relations {
		if v.rel == Parent && v.from.name == name {
			res = append(res, r.relations[i].to)
		}
	}
	return res
}

// high level module
type Research struct {
	// break dependency inversion principle
	// relationships Relationships
	// abstraction!
	browser RelationshipBrowser
}

// func (r *Research) Investigate() {
// 	relations := r.relationships.relations

// 	for _, rel := range relations {
// 		if rel.from.name == "John" && rel.rel == Parent {
// 			fmt.Println("John has a child called ", rel.to.name)
// 		}
// 	}
// }

func (r *Research) InvestigateAbs() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called ", p.name)
	}
}

func main() {
	j := Person{name: "John"}
	c := Person{"Chris"}
	c2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&j, &c)
	relationships.AddParentAndChild(&j, &c2)

	// r := Research{relationships}
	r := Research{&relationships}
	r.InvestigateAbs()
}
