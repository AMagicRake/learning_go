package main

import "fmt"

// A new type called 'deck' that is a slice of string
type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
