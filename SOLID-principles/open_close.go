package main

import "fmt"

// OCP
// open for extension, closed for modification
// Specification... enterprise manner

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	// filtering..

}

func (f *Filter) FilterByColor(prods []Product, color Color) []*Product {
	res := make([]*Product, 0)

	for i, v := range prods {
		if v.color == color {
			res = append(res, &prods[i])
		}
	}

	return res
}

// specification pattern

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (c SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == c.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(prods []Product, spec Specification) []*Product {
	res := make([]*Product, 0)

	for i, v := range prods {
		if spec.IsSatisfied(&v) {
			res = append(res, &prods[i])
		}
	}

	return res
}

// can also implement ANDSpecification or ORSpecification for filtering..

// func main() {
// 	apple := Product{"Apple", green, small}
// 	tree := Product{"Tree", green, large}
// 	house := Product{"House", red, large}

// 	prods := []Product{apple, tree, house}
// 	f := Filter{}

// 	fmt.Printf("Green prods (old) : \n")
// 	for _, v := range f.FilterByColor(prods, green) {
// 		fmt.Printf(" - %s is green \n", v.name)
// 	}

// 	fmt.Printf("Green prods (new) : \n")
// 	greenSpec := ColorSpecification{green}
// 	bf := BetterFilter{}
// 	for _, v := range bf.Filter(prods, greenSpec) {
// 		fmt.Printf(" - %s is green \n", v.name)
// 	}
// }
