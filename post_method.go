package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	todo := Todo{
		UserId:    1,
		Title:     "delectus aut autem",
		Completed: false,
	}
	// converting struct to json  (Marshalling)
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error in marshalling : ", err)
		return
	}

	//convert json data to string
	jsonString := string(jsonData)

	//posting

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	jsonReader := strings.NewReader(jsonString)

	res1, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error in posting data : ", err)
		return
	}
	defer res1.Body.Close()

	// converting response to readable format
	data, _ := ioutil.ReadAll(res1.Body)
	fmt.Println(string(data))
	fmt.Println("Response status", res1.Status)

}
