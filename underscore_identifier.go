package main

import "fmt"

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Cannot divide by zero"
	}
	return a / b, ""
}

func main() {
	fmt.Println("Hello, World!")
	ans, err := divide(10, 0)
	fmt.Println("ans : ", ans)
	fmt.Println("err : ", err)

}
