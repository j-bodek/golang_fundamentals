package main

import "fmt"

// define contact info struct
type contactInfo struct {
	email   string
	zipCode int
}

// define person struct
type person struct {
	firstName string // define firstName field
	lastName  string // define lastName field
	contact   contactInfo
}

func main() {
	// declare variable of type person
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@email.com",
			zipCode: 12345,
		},
	}

	// declare pointer to alex variable (& - returns memory address of the alex variable)
	// alexPointer := &alex

	// update alex firstName
	alex.updateName("John") // update lastName field value
	alex.print()            // print person fields with it values
	fmt.Printf("Person first name is: %s\n", alex.firstName)
	fmt.Printf("Person last name is: %s\n", alex.lastName)
}

// function used to update firstName field, add person pointer as receiver
// *person - type pointer that point to person
func (pointerToPerson *person) updateName(newFirstName string) {
	// update variable that is pointed by pointer
	// * - returns value of variable stored in specified memory address
	(*pointerToPerson).firstName = newFirstName
}

// function used to print variable of type person
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
