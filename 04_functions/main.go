package main

import "fmt"

// should declare param type, second is return value

func sayHi (name string) string {
		return "Hello " + name
}

func getTotal (num1, num2 int) int {
	return num1 + num2
}

func main () {
	fmt.Println(sayHi("hector"))
	fmt.Println(getTotal(100, 50))
}