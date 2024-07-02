package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{
		Name:    "divyanshu",
		Age:     23,
		IsAdult: true,
	}
	fmt.Println(person)

	// converting struct to json  (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error in marshalling : ", err)
	}
	fmt.Println("json data : ", jsonData)
	fmt.Println("json data in string : ", string(jsonData))

	//decoding json data to struct (Unmarshalling)

	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("error in unmarshalling : ", err)
	}
	fmt.Println("Person Data : ", personData)

}
