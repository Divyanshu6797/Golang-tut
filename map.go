package main

import "fmt"

// map is unordered collection of key value pairs

func main() {
	studentGrades := make(map[string]int)
	studentGrades["John"] = 92
	studentGrades["Jane"] = 100
	studentGrades["Doe"] = 90

	fmt.Println("John's grade:", studentGrades["John"])

	// delete an entry
	delete(studentGrades, "John")

	// access using range
	for key, value := range studentGrades {
		fmt.Println("key:", key, "value:", value)
	}

	// checking if a key is present
	grade, exists := studentGrades["John"]
	if exists {
		fmt.Println("John's grade:", grade)
	} else {
		fmt.Println("John's grade: not found")
	}

	// another map with string values
	person := map[string]string{
		"name": "John",
		"age":  "25",
		"city": "New York",
	}

	for key, value := range person {
		fmt.Println("key:", key, "value:", value)
	}
}
