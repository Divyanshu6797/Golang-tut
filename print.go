package main

import "fmt"

func main() {

	age := 25
	name := "divyanshu"
	heigth := 5.85454

	fmt.Println("name : ", name, " age : ", age, " heigth : ", heigth) // gives a line break after printing

	fmt.Printf("name : %s age : %d heigth : %0.3f\n", name, age, heigth)

	fmt.Printf("type of name : %T\n", name)
}
