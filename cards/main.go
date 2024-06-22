package main

func main() {

	cardDeck := newDeck()
	cardDeck.shuffle()
	cardDeck.print()
	// cardDeck.saveToFile("deckState.dck")

	// readDeck := newDeckFromFile("deckState.dck")
	// readDeck.print()
}
