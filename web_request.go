package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error : ", err)
	}

	defer res.Body.Close()
	fmt.Printf("type of Response :%T", res)

	// reading content of response

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error : ", err)

		return
	}
	fmt.Println(string(data))

}
