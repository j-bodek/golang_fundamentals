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
