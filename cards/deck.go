package main

import "fmt"

// Create new deck type
// which is a slice of strings
type deck []string // tell that deck 'inherit' from string

// declare new function that belongs to deck type
// deck is 'receiver' so any variable of type deck has access to print method
// d - reference to object, deck - type of object
func (d deck) print() { // (d deck) tells go that print function belongs to deck type
	// iterate through cards elements
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	// create list of suits and values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// create suit, values combinations
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	// return cards
	return cards
}

// declare deal function that takes two arguments:
// deck, handSize and return two decks
func deal(d deck, handSize int) (deck, deck) {
	// split d value into two parts
	return d[:handSize], d[handSize:]
}
