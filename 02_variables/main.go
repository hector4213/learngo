package main
import "fmt"


func main() {
	// MAIN TYPES

	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint32 uint 64 iuntptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//Using var
  var name string  = "Hector"
	var age int = 99
	var isCool = true

	// Short inferred syntax
	nickName := "Hectorious"
	birthYear := "1935"

	city, petName := "Toronto", "Meowboy"


	fmt.Println(name, age, isCool)
	fmt.Println(nickName, birthYear)
	fmt.Println(city, petName)
	fmt.Printf("%T\n", isCool)
}