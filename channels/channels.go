package channels

import (
	"fmt"
	"time"
)

func Execute() {
	fmt.Println("Start Example ----> 1.9.2")
	example()
	fmt.Println("End Example ----> 1.9.2")
}

func sum(channel chan int) int {
	s := 0

	for v := range channel {
		s = s + v
		fmt.Printf("Parcial value %d\n", s)
	}

	return s
}

func example() {
	channel := make(chan int, 5)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			time.Sleep(time.Second)
		}
		close(channel)
	}()

	fmt.Println(sum(channel))
}
