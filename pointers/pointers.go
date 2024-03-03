package pointers

import "fmt"

func Execute() {
	fmt.Println("Start Example ----> 1.5.3")
	example()
	fmt.Println("End Example ----> 1.5.3")
}

/*
here we are changing the value with a pointers
output 5
*/
func example() {
	// set 8 to i var
	i := 8

	// declare pointer
	var pi *int

	// set pointer i to pi
	pi = &i

	fmt.Printf("Prev to set value %v. \n", i)

	// set new value to the pointer
	// this pointer is the same that have the var i
	*pi = 5

	fmt.Printf("Post to set new value %v. \n", i)
}
