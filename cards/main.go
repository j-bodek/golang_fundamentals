package main

func main() {
	// declare variable of string type
	// basic go types: bool, string, int, float64
	// card := "Ace of Spades" // shortcut of below line, this declaration works with any types
	// var card string = "Ace of Spades"

	// declare slice of strings
	// cards := newDeck()
	// cards.saveToFile("deck")
	cards := newDeckFromFile("deck")
	cards.print()
}

func newCard() string { // after () declare that function return variable type string
	return "Five of Diamonds"
}
