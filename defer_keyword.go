package main

import "fmt"

func main() {

	fmt.Println("starting of the program ")
	defer fmt.Println("middle of the program ")  // executed at last of main function
	defer fmt.Println("middle of the program 1") // executed at last of main function
	fmt.Println("ending of the program ")
}

// where there are more than one defer statements, they are executed in LIFO order
