package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach
type EmployeeFactory struct {
	position     string
	annualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.position, f.annualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	// dF, mF are functions... take name => employee
	// you cannot change some parameters, you have to recreate it.
	dF := NewEmployeeFactory("developer", 80000)
	mF := NewEmployeeFactory("manager", 100000)

	d := dF("Sid")
	m := mF("Ram")

	fmt.Println(d, m)

	// bF is a struct, so it has receiver func like Create()
	// you can change some parameters, you don't have to recreate it.
	bF := NewEmployeeFactory2("CEO", 1000000)
	b := bF.Create("Sid")
	fmt.Println(b)
}
