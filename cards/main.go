package main

import "fmt"

func main() {
	card := newCard()
	cardDeck := []string{"Ace Of Spades", "Two of Spades"}
	fmt.Println(card)
	fmt.Println(cardDeck)
	cardDeck = append(cardDeck, "Three of Spades")
	for _, card := range cardDeck {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
