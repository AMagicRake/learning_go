package main

func main() {

	cardDeck := newDeck()
	var hand deck

	cardDeck.print()

	cardDeck, hand = deal(cardDeck, 4)
	hand.print()

}
