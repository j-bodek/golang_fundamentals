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

	// loop through every link and try to open it
	for _, link := range websites {
		fmt.Printf("Trying to open: '%s'\n", link)
		checkLink(link)
	}
}

// function used to check if link returns error after get request
func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("Something wrong with '%s' %s\n\n", link, err)
	} else {
		fmt.Printf("'%s' is up\n\n", link)
	}
}
