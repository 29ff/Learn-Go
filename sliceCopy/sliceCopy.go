package main

import "fmt"

func main() {
	var number []int
	// create a new slice
	number2 := make([]int, 15)

	// copy the original slice to the new slice
	copy(number2, number)
	fmt.Println(number)
}
