package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{height: 10.0, base: 10.0}
	s := square{sideLength: 5.0}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (sq square) getArea() float64 {
	return math.Pow(sq.sideLength, 2)
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.height * tr.base
}
