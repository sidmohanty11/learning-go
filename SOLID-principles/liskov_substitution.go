package main

import "fmt"

// Liskov Substitution

type Sized interface {
	GetWidth() int
	SetWidth(w int)
	GetHeight() int
	SetHeight(h int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	w := sized.GetWidth()
	sized.SetHeight(10)

	expectedArea := 10 * w
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Print("Expected area = ", expectedArea, " and actual area = ", actualArea)
}

type Square struct {
	Rectangle
}

// will violate liskov_substitution principle
func NewSq(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// for get around
type SquareFixed struct {
	size int
}

// now it will construct a new rectangle with width and height set to the size of the Square
func (s *SquareFixed) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

// func main() {
// 	rc := &Rectangle{
// 		width:  20,
// 		height: 14,
// 	}
// 	UseIt(rc)

// 	sq := &SquareFixed{
// 		size: 10,
// 	}
// 	s := sq.Rectangle()
// 	UseIt(&s)
// }
