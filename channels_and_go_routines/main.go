package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{
		"https://google.com",
		"https://golang.org",
		"https://facebook.com",
		"https://amazon.com",
		"https://stackoverflow.com",
	}

	// create new channel to allow communication between
	// child and main routines
	c := make(chan string) // channel that can be used to send string

	// loop through every link and try to open it
	for _, link := range websites {
		// create new go routine (checkLink will be executed in new
		// go routine to prevent main routine when waiting for response)
		go checkLink(link, c) // pass channel to checkLink
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}
}

// function used to check if link returns error after get request
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("Something wrong with '%s' %s\n\n", link, err)
		c <- "something went wrong"
	} else {
		fmt.Printf("'%s' is up\n\n", link)
		c <- "website is up"
	}
}
