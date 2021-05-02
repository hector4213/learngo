package main

import "fmt"


func main () {
	// Arrays are of fixed length in Go
  var fruitArr [2]string

	//Assign Vals
	fruitArr[0] = "Mango"
	fruitArr[1] = "PineApple"

	// Declare and assign in one go
	numArr := [2]int{69,69}

	// Slices are kind of like arrays, but they have no fixed length
	characters := []string{"Goku", "Vegeta", "Krillin", "Yamcha"}

	// Get the length
	howManyBois := len(characters)

	// Get a range 
	justVegeta := characters[1:2]
	onlySaiyans := characters[0:2]

	fmt.Println(fruitArr, numArr)
	fmt.Println(characters, howManyBois, justVegeta, onlySaiyans)
}