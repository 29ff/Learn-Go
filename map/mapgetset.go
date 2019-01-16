package main

import "fmt"

func main() {
	var m map[string]int

	// adding key/value
	m["clearity"] = 2
	m["simplicity"] = 3

	// printing the values
	fmt.Println(m["clearity"])
	fmt.Println(m["simplicity"])
}
