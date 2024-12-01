package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck' that is slice of strings

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	cards := deck{}
	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, "\n")
}

func (d deck) saveToFile(filename string) error {
	fmt.Printf("Writing to the file: %s\n", filename)
	return os.WriteFile(filename, []byte(d.toString()), 0622)

}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {

	}
	return deck(strings.Split(string(bs), "\n")), err
}

func (d deck) shuffle() deck {
	for i := range d {
		randomIndex := rand.Intn(len(d))
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
	return d
}
