package defers

import "fmt"

func Execute() {
	fmt.Println("Start Example ----> 1.6.5")
	example()
	fmt.Println("End Example ----> 1.6.5")
}

/*
the defers are execute after the function, and if you declare multiple defer are stacked
and execute in LIFO order
output -> A, C, D, B
*/
func example() {
	fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("C")
	defer fmt.Println("D")
}
