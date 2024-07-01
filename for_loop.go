// go doesnt have while loop or do while loop
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i : ", i)
	}

	// infinite loop, can be used like while loop
	counter := 0
	for {
		fmt.Println("Hello, World!")
		counter++
		if counter == 10 {
			break // continue is also here
		}
	}

	// range keyword is used to iterate over items of an array, slice, string, map or channel
	number := []int{0, 323, 334, 54, 65}
	for index, value := range number {
		fmt.Println("index : ", index, "value : ", value)
	}

	data := "Hello world"

	for index, value := range data {
		fmt.Println("index : ", index, "value : ", string(value))
	}

}
