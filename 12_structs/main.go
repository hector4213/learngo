package main

import (
	"fmt"
	"strconv"
)

// Define a struct

type User struct {

	email, username, password, gender string
	age int
}

// Greeting method (value receiver)

func (u User) greet() string {
	return "Happy Birthday! " + u.username + " you are " + strconv.Itoa(u.age)
}

// has Birthday method (pointer receiver) (change something)

func (u *User) hasBirthday() {
	u.age++
}

// gotMarried (pointer receiver)

func (u *User) gotMarried(newLastName string) {
	if u.gender == "female" {
		u.username = "Mrs." + u.username + " " + newLastName
	} else {
		return
	}
}

func main() {
	// Initialize a person using struct
	newUser := User{email: "test@test.com", username: "testy", password: "hunter2", gender: "male", age:99}

	// Shorthand
	anothaOne := User{"me@mail.com", "kawahi_leonard", "hunter3", "female", 12}

	// Structs have two kind of methods, value receivers and pointer receivers
	anothaOne.gotMarried("Lopez")
	anothaOne.hasBirthday()
	fmt.Println(anothaOne.greet())
	fmt.Println(newUser, anothaOne)


}
