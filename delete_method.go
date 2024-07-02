package main

import (
	"fmt"
	"net/http"
)

func main() {

	// deleting todo with id 1

	// creating delete request
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		println("error in creating request : ", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println("error in deleting data : ", err)
	}
	defer res.Body.Close()
	fmt.Println("Response status : ", res.Status)

}
