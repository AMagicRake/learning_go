package main

import "fmt"

func main() {

	cardDeck := newDeck()
	fmt.Println(cardDeck.toString())
	cardDeck.saveToFile("deckState.dck")

}
