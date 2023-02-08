package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 99999) // initialize empty byte slice with 99999 available element spaces
	// read response body to data slice
	resp.Body.Read(data)
	fmt.Println(string(data))
}
