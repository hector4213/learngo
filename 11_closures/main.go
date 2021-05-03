package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func multiplier () func(int) int {
	product:= 1
	return func (x int) int {
		product *= x
		return product
	}
}

func main() {

	sum := adder()
	
	products := multiplier()

	for i := 1; i < 10; i++ {
		fmt.Println(products(i))
	}

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
