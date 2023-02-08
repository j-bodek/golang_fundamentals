package main

import "fmt"

// create bot interface (every type that implement getGreeting method
// will be 'included' as type bot)
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// method used to generate english greetings
func (eb englishBot) getGreeting() string {
	// custom logic for generating an english greeting
	return "Hi There!"
}

// method used to generate spanish greetings
func (sb spanishBot) getGreeting() string {
	// custom logic for generating an spanish greeting
	return "Hola!"
}
