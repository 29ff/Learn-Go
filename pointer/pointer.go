package main

import "fmt"

func main() {
	var x *int
	var y int

	y = 0
	x = &y

	fmt.Println(x)  // => 0xc000068008
	fmt.Println(&x) // => 0xc00007a018

	*x = 1

	fmt.Println(*x) // => 1
	fmt.Println(y)  // => 1
}
