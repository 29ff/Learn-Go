package main

import "fmt"

func main() {
	// initialize a slice with 4 len and values
	var number2 = []int{1, 2, 3, 4}
	fmt.Println(number2) // => [1, 2, 3, 4]

	// create sub slices
	slice1 := number2[2:]
	fmt.Println(slice1) // => [3, 4]

	slice2 := number2[:3]
	fmt.Println(slice2) // => [1, 2, 3]

	slice3 := number2[1:4]
	fmt.Println(slice3) // => [2, 3, 4]

	slice4 := number2[2:4]
	fmt.Println(slice4) // => [3, 4]

	slice5 := number2[number2[0]:number2[len(number2)-1]]
	fmt.Println(slice5) // => [2, 3, 4]
}
