package main

func main() {
	// declare variable of string type
	// basic go types: bool, string, int, float64
	// card := newCard() // shortcut of below line, this declaration works with any types
	// var card string = "Ace of Spades"

	// declare slice of strings
	cards := deck{"Ace of Spades", newCard()}
	// append new value to cards (notice that append returns new slice)
	cards = append(cards, "Six of Spades")
	// print function is declared in deck.go file
	cards.print()
}

func newCard() string { // after () declare that function return variable type string
	return "Five of Diamonds"
}
