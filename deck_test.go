package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace Of Heart" {
		t.Errorf("Expected deck[0] \"Ace Of Heart\", but got %v", d[0])
	}

	if d[len(d)-1] != "King Of Clubs" {
		t.Errorf("Expected last card \"King Of Clubs\", but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	dFromFile := newDeckFromFile("_decktesting")
	if len(dFromFile) != 52 {
		t.Errorf("Ecpected 52 cards in deck, but got %v", len(dFromFile))
	}

	os.Remove("_decktesting")
}
