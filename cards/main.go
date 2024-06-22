package main

func main() {

	cardDeck := newDeck()
	cardDeck.saveToFile("deckState.dck")

	readDeck := newDeckFromFile("deckState.dck")
	readDeck.print()
}
