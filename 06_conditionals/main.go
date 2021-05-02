package main

import "fmt"


func main () {
	x := 4 
	y := 25


	// If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Else if

	color := "pink"

	if color == "orange" {
		fmt.Println("color is red")
	} else if color == "pink" {
		fmt.Println("color is pink")
	} else {
		fmt.Println("What kind of color is this")
	}

	// Switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	case "pink":
		fmt.Println("color is pink")
	default:
		fmt.Println("what color is this???")
	}

}