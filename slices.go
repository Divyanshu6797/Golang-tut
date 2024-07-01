// slices have a dynamic size
package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Printf("type of numbers : %T\n", numbers)

	numbers = append(numbers, 45)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	// make function to create a slice with length and capacity
	numbers1 := make([]int, 5, 10) // length = 5, capacity = 10
	fmt.Println(len(numbers1))
	numbers1 = append(numbers1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(numbers1)
	fmt.Println(len(numbers1))

	fmt.Println(cap(numbers1))

	//slice with length and capacity 0
	numbers2 := make([]int, 0) // empty slice
	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))

}
