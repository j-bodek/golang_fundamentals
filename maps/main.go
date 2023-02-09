package main

import "fmt"

func main() {
	// STRUCT VS MAPS

	// Map :
	// - All keys must be the same type
	// - All values must be the same type
	// - Keys are indexed we can iterate over them
	// - Reference Type
	// - Use to represent a collection of related properties

	// Struct:
	// - Values can be of different type
	// - Keys don't support indexing (we cannot iterate over them)
	// - Value Type
	// - Use to represent a "thing" with a lot of different properties

	// colors := make(map[int]string)
	// colors[10] = "red"
	// delete(colors, 10)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf723",
	}

	// map iteration
	printMap(colors)

	// fmt.Println(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("%s -≥ %s\n", key, value)
	}
}
