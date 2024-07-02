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
		Title:     "this is new updated data",
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

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	jsonReader := strings.NewReader(jsonString)

	// making a request
	req1, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)

	if err != nil {
		fmt.Println("error in creating request : ", err)
		return
	}
	req1.Header.Set("Content-Type", "application/json")

	// sending request
	client := http.Client{}
	res2, err := client.Do(req1)
	if err != nil {
		fmt.Println("error in posting data : ", err)
		return
	}

	defer res2.Body.Close()

	// converting response to readable format
	data, _ := ioutil.ReadAll(res2.Body)
	fmt.Println(string(data))
	fmt.Println("Response status", res2.Status)

}
