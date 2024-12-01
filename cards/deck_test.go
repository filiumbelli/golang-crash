package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected 52, got %d", len(deck))
	}

	firstElement := deck[0]
	if firstElement != "Ace of Spades" {
		t.Errorf("Expected first element Ace of Spades, got %s", firstElement)
	}
	if deck[len(deck)-1] != "Ten of Clubs" {
		t.Errorf("Expected last element Ten of Clubs, got %s", deck[len(deck)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	d1, _ := newDeck().deal(5)
	filename := "test.txt"
	d1.saveToFile(filename)

	readDeck, e := newDeckFromFile(filename)

	if e != nil {
		t.Errorf("Expected errror to be nil, got %v", e)
	}
	if len(readDeck) != 5 {
		t.Errorf("Expected deck read by file with the length of 5, got %d", len(readDeck))
	}
	for i := range readDeck {
		if readDeck[i] != d1[i] {
			t.Errorf("Expected deck read(%s) to be equal to the deck created(%s) but got different at index %d. ", readDeck[i], d1[i], i)
		}
	}
	os.Remove(filename)
}
func TestDealDeck(t *testing.T) {
	deck := newDeck()
	d1, d2 := deck.deal(5)

	if len(d1) != 5 {
		t.Errorf("Expected deck length to be 5, got %d", len(d1))
	}
	if len(d2) != 47 {
		t.Errorf("Expcted deck length to be 47, got %d", len(d2))
	}

}
