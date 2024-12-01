package main

import (
	"fmt"
	"os"
)

var deckSize int = 50

func main() {
	deal1, _ := newDeck().deal(3)
	deal1.saveToFile("deck1.txt")
	d1, _ := newDeck().shuffle().deal(10)
	d1.print()
	newDeck, err := newDeckFromFile("deck1.txt")
	if err != nil {
		fmt.Println("ERROR: Could not load deck the deck from file " + err.Error())
		os.Exit(1)
	} else {
		newDeck.shuffle()
	}
}
