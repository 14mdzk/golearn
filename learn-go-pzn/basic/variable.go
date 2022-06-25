package main

import "fmt"

func main() {
	// var (name) (data type *optional)
	// var is used for declare a variable
	// name is the name for variable
	// data type is declare what kind of data type stored in variable
	// each variable is locked based on its data type
	// variable with string data type can't be used for storing number data etc
	var name string

	// or variable can be declared using
	// name := value
	
	// and you can multi declare by doing so
	// var ( name1 = value1 name2 = value2 ...  )
	var (
		firstName = "Your"
		lastName = "Name"
	)

	const phi = 3.14
	// const variable can't be reassigned or it's value can't be changed

	// unsued variable can lead to syntax error

	name = "My Name"

	fmt.Println(name)
	fmt.Println("First Name: ", firstName)
	fmt.Println("Last Name: " , lastName)
	fmt.Println("Phi: ", phi)
}
