package functions

import "fmt"

func Execute() {
	fmt.Println("Start Example ----> 1.7.1")
	example()
	fmt.Println("End Example ----> 1.7.1")
}

func fibonacci() func() int {
	f1 := 1
	f2 := 2

	return func() int {
		f2, f1 = f2+f1, f2
		return f1
	}
}

func example() {
	fb := fibonacci()

	for i := 0; i < 5; i++ {
		fmt.Printf("Fibonacci %d: %d - ", i, fb())
	}
}
