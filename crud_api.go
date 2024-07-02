package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		println("error in fetching data : ", err)
	}
	defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error in reading data : ", err)
	// }
	// println(string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error in decoding data : ", err)
	}
	fmt.Println("Todo : ", todo)

	var todoArr []Todo
	data, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("error in fetching data : ", err)
	}
	defer data.Body.Close()

	err = json.NewDecoder(data.Body).Decode(&todoArr)
	if err != nil {
		fmt.Println("error in decoding data : ", err)
	}
	fmt.Println("Todo Array : ", todoArr)

}
