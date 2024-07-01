package main

import "fmt"

func main() {
	var name string = "i am string"
	var version = 86 // datatype is decided in copile time
	var dimension float64 = 2.5
	dimension = 66.5

	const pi = 3.14 // const value cannot be changed

	fmt.Println(name)
	fmt.Println(version)
	fmt.Println(dimension)
	fmt.Println(pi)

	person := "iam person" // shorthand declaration
	fmt.Println(person)

	var Blah = "i am public"  // can be used in another package
	var blah = "i am private" // cannot be used in another package

}
