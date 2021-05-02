package main

import "fmt"

func main () {
	ids := []int{66,44,55,66,55,33,42,59,75,54,34}

	for i, id := range ids {
		fmt.Printf("im at index %d and the id is %d\n", i, id)
	}

	// Not using index, use '_' to omit unused vars

	for _, id := range ids {
		fmt.Printf("I am number %d\n", id)
	}

	// Sum ids

	sum := 0
	for _, id := range ids {
		sum+= id
	}
	fmt.Printf("The total is %d and type is %T\n", sum, sum)

	// Range with Map
	person := map[string]string{"name": "John", "city":"New York", "occupation": "painter"}

	for key, value := range person {
		fmt.Printf("%s  %s\n", key, value)
	}

	for key:= range person {
		fmt.Println("The key is " + key) 
	}


}