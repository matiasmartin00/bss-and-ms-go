package slices

import "fmt"

func Execute() {
	fmt.Println("Start Example ----> 1.5.6")
	example()
	fmt.Println("End Example ----> 1.5.6")
}

func example() {
	// declare array int with size 5
	s1 := make([]int, 5)

	// declare array int with size 0, capacity 5
	s2 := make([]int, 0, 5)

	s1 = append(s1, 1)
	s2 = append(s2, 2)

	// how we added an element, the length is 6 and the capacity is 10 (double)
	fmt.Printf("S1: Length %d, Capacity %d. \n", len(s1), cap(s1))

	// how we added an element, in this case the length is 1 and the capacity is 5
	fmt.Printf("S2: Length %d, Capacity %d. \n", len(s2), cap(s2))
}
