package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		log.Fatal(err)
	}

	// data := make([]byte, 99999) // initialize empty byte slice with 99999 available element spaces
	// // read response body to data slice
	// resp.Body.Read(data)
	// fmt.Println(string(data))

	// Alternative
	io.Copy(os.Stdout, resp.Body)
}

// create Write method to implement writer interface
func (lw logWriter) Write(bs []byte) (int, error) {
	// print byte slice
	fmt.Println(string(bs))
	return len(bs), nil
}
