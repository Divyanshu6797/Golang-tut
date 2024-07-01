package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,grapes,apple"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	count := strings.Count(data, "apple")
	fmt.Println(count)

	str := "      hemlo i am cheems     "
	trimmed := strings.TrimSpace(str) // remove whitespaces from start and end
	fmt.Println(trimmed)

	str1 := "hello"
	str2 := "world"
	joined := strings.Join([]string{str1, str2}, " <i am separator> ")
	fmt.Println(joined)

}
