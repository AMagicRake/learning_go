package main

import (
	"os"
	"testing"
)

// t.Errorf("Expected deck length of 16, but got %v", len(d))
// test card count
func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(d))
	}
}

//test card uniqueness

func TestDeal(t *testing.T) {
	handSize := 5
	deckSize := len(newDeck()) - handSize
	d, h := deal(newDeck(), handSize)
	if len(h) != handSize {
		t.Errorf("Expected handsize of 5, but got %v", len(h))
	}
	if len(d) != (deckSize) {
		t.Errorf("Expected decksize of %v, but got %v", (deckSize), len(d))
	}
}

func TestSaveAndLoadDeck(t *testing.T) {
	fileName := "deckState.test.dck"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)
	if len(deck) != len(loadedDeck) {
		t.Errorf("Expected decksize of %v, but got %v", len(deck), len(loadedDeck))
	}
	for i := range deck {
		if deck[i] != loadedDeck[i] {
			t.Errorf("The value of loadedDeck at index %v is not the same as the written deck at the same index.", i)
		}
	}
	os.Remove(fileName)
}
