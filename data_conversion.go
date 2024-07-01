package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 2
	fmt.Printf("type of num : %T\n", num)

	var data float64 = float64(num)
	fmt.Printf("type of data : %T\n", data)

	hi := 5
	str := strconv.Itoa(hi)
	fmt.Printf("type of str : %T\n", str)

	number := "10"
	num1, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("type of num1 : %T\n", num1)

	str1 := "10.5"
	num2, err := strconv.ParseFloat(str1, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("type of num2 : %T\n", num2)

}
