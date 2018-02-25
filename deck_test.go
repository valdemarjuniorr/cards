package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Exptected deck length of 16, but got %v", len(cards))
	}
}

func TestFirstDeckElement(t *testing.T) {
	cards := newDeck()

	if cards[0] != "Aces of Spades" {
		t.Errorf("Expecting Aces of Spades, but got %v", cards[0])
	}
}

func TestLastDeckElement(t *testing.T) {
	cards := newDeck()
	lastCard := cards[len(cards)-1]

	if lastCard != "Four of Clubs" {
		t.Errorf("Expecting Four of Clubs, but got %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	cards := newDeck()

	cards.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expecting 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")

}
