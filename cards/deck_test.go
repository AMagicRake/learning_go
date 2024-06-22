package main

import "testing"

//t.Errorf("Expected deck length of 16, but got %v", len(d))
//test card count
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
