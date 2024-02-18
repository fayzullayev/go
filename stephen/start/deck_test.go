package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of length of Ace of Spades, but got %s", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of length of Four of Clubs, but got %s", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()

	deck.SaveToFile("_decktesting")

	loadedDeck := ReadFromFile("_decktesting")

	if len(loadedDeck) != 17 {
		t.Errorf("Expected deck length of 16, but got %d", len(deck))
	}

	os.Remove("_decktesting")
}
