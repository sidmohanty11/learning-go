package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//create a new type of deck -> slice of strings
//->type deck extends all props of slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// reciever funcs -> before the name of the function
// d => ref to actual var or -> cards(var) in main.go
// deck => type

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//slice[startingIndexIncluding: endingIndexWithoutIncluding]
// 0 -> 2 (to/upto)
// [0:2] = [:2]
//if funcName(...,...) -> arguments!
//returning two separate values
//deal(cards, handSize) -> 2 slices
//example- handSize -> 26
//give two slices -> 26cards, 26cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) deckToString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}
