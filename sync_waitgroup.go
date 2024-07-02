package main

import (
	"fmt"
	"sync"
)

func Worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker started : ", i)
	fmt.Println("Worker completed : ", i)

}
func main() {

	var wg sync.WaitGroup
	// start 3 worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("workers completed")

}
