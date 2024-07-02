// goroutine allows us to run a function concurrently with other functions.
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond) // simulating some work
	fmt.Println("Hello again")
}
func sayWorld() {
	fmt.Println("World")
}

func main() {
	go sayHello()
	sayWorld()

	time.Sleep(3000 * time.Millisecond)

}
