package main

import "fmt"

func main() {

	a := 5

	// Pointer to memory address 
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // * int

	// Use * to read val from address
	fmt.Println(*b)

	// Change value with pointer
	*b = 100
	// Pointers are used to pass data stored at an address, if you choose to pass the address instead of the data itself
	// it can lead to better performance?
	// Everything in Go is pass by value
	fmt.Println(a)
}