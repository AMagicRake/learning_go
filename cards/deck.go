package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// A new type called 'deck' that is a slice of string
type deck []string

func newDeck() deck {
	d := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}

	return d
}

func newDeckFromFile(fileName string) deck {
	readFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	d := strings.Split(string(readFile), ",")
	return d
}

func deal(d deck, handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) {
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
