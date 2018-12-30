package main

import "fmt"

func main() {
	var m map[string]int

	m['a'] = 2
	m['b'] = 3

	fmt.Println(m['a'])
	fmt.Println(m['b'])
}