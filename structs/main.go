package main

import "fmt"

// define person struct
type person struct {
	firstName string // define firstName field
	lastName  string // define lastName field
}

func main() {
	// declare variable of type person
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("Person first name is: %s\n", alex.firstName)
	fmt.Printf("Person last name is: %s\n", alex.lastName)
}
