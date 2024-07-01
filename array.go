package main

import "fmt"

func main() {

	var name [3]string
	name[0] = "divyanshu"
	name[1] = "kumar"
	name[2] = "singh"

	fmt.Println(name)
	fmt.Println(name[0])

	// array inialization

	var number = [3]int{1, 2, 3}
	fmt.Println(number)

	fmt.Println("length of number array : ", len(number))

	// by default all elements are initialized to 0
	var number1 [5]int
	fmt.Println(number1)

	// 2D array
	var matrix [2][2]int
	matrix[0][0] = 1

	var price [5]string
	fmt.Printf("%q", price) // prints quoted string

}
