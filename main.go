package main

func main() {
	// cards := newDeck()
	// hand, cardsRemaning := deal(cards, 5)
	// hand.print()
	// cardsRemaning.print()

	cards := newDeck()
	cards.shuffle()
	// cards.saveToFile("my_cards")
	// cards = newDeckFromFile("my_cards")

	cards.print()
}
