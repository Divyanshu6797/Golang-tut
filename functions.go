package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func newAdd(a int, b int) (result int) {
	result = a + b
	return
}

func main() {
	ans := add(1, 20)
	fmt.Println("ans : ", ans)
	ans1 := newAdd(1, 20)
	fmt.Println("ans1 : ", ans1)

}
