package main

import "fmt"

func main() {
	card := newCard()
	cardDeck := newDeck()
	fmt.Println(card)
	fmt.Println(cardDeck)
	cardDeck = append(cardDeck, "Three of Spades")

	cardDeck.print()

}

func newCard() string {
	return "Five of Diamonds"
}
