package main

import "fmt"

func zeroToNSlice(n int) []int {
	numbers := []int{}
	for i := 0; i < n+1; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func main() {
	numbers := zeroToNSlice(10)
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%d is even number\n", number)
		} else {
			fmt.Printf("%d is odd number\n", number)
		}
	}
}
