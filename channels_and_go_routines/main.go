package main

import (
	"fmt"
	"net/http"
	"time"
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

	// listen for messages from channels (similar like asyncio look in python)
	for link := range c {
		// run checkLink as function literal (function literal == python lambda function)
		go func(l string) {
			time.Sleep(2 * time.Second)
			checkLink(l, c)
		}(link)
	}
}

// function used to check if link returns error after get request
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Something wrong with '%s' %s\n\n", link, err)
		c <- link
	} else {
		fmt.Printf("'%s' is up\n\n", link)
		c <- link
	}
}
