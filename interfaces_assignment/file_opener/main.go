package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You have to provide filename")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)
}
