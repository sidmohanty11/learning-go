package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i] //one line swap
	}
}
