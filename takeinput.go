package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var name string
	fmt.Scan(&name) // reades only till white space
	fmt.Println("name : ", name)

	reader := bufio.NewReader(os.Stdin)
	name1, _ := reader.ReadString('\n') // reads till new line
	fmt.Println("name : ", name1)

}
