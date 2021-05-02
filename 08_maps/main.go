package main

import "fmt"


func main () {
	// Make a map (key value pairs)

	emails := make(map[string]string)

	// Assign kv
	emails["Jose"] = "tooSpicy@hotsauce.com"
	emails["Consuela"] = "MrsClean4u@casita.com"

	// Declare map AND add key value pairs

	secretEmails := map[string]string{"Jose" : "NoTellMyWife33@mail.com", "Consuela": "imUnlIcense@tito.com"}

	fmt.Println(emails)

	// Get length of map
	
	fmt.Println(len(emails))
	fmt.Println(emails["Consuela"])
	fmt.Println(emails)
	fmt.Println(secretEmails)

}