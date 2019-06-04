package main

import "fmt"

// Create a new *type* of 'deck
// which will be a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Clubs", "Diamonds", "Spades"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	// nested iteration
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func (d deck) print()  --> the `d deck` is the receiver - Any variable of type "deck" now gets access to the print method. The receiver sets

// Selecting a range of slices syntactically: `exampleSlice[startIndexIncluding: uptoNotIncluding]`
// eg selecting the first two elements of a slice of 4 elements would look like `exampleSlice[0:2]`
