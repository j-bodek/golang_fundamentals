package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Create new deck type
// which is a slice of strings
type deck []string // tell that deck 'inherit' from string

// DECK METHODS

// declare new function that belongs to deck type
// deck is 'receiver' so any variable of type deck has access to print method
// d - reference to object, deck - type of object
func (d deck) print() { // (d deck) tells go that print function belongs to deck type
	// iterate through cards elements
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function that converts deck type variable into string
func (d deck) toString() string {
	// first argument has to be slice of strings
	return strings.Join([]string(d), ",")
}

// function used to save deck to file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// this function is used to shuffle deck
func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i int, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// HELPER FUNCTIONS

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

// function that create new deck from file
func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)

	// check if no error (nil is like None in python)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program

		// Option #2:
		// log.Fatal works like Print followed by os.Exit
		log.Fatal(err)
	}

	return deck(strings.Split(string(data), ","))
}

// declare deal function that takes two arguments:
// deck, handSize and return two decks
func deal(d deck, handSize int) (deck, deck) {
	// split d value into two parts
	return d[:handSize], d[handSize:]
}
