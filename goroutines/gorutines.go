package goroutines

import (
	"fmt"
	"sync"
)

func Execute() {
	fmt.Println("Start Example ----> 1.9.1")
	example()
	fmt.Println("End Example ----> 1.9.1")
}

func example() {
	counter := 0
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock() // without mutex we have a race condition
			counter = counter + 1
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter result -> %d\n", counter)
}
