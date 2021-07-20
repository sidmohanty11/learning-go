package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //t -> test handler
	d := NewDeck() //can't import function from another file

	if len(d) != 52 {
		t.Errorf("Expected deck len of 52 but got, %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := NewDeck() //can't import function from another file
	deck.saveToFile("_decktesting")

	loadedDeck := NewDeckFromFile("_decktesting") //can't import function from another file

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
