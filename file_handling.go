package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// creating a file
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Println("file created successfully")

	content := "hi am learning go"

	// writing content to the file
	byte1, errors := io.WriteString(file, content)
	if errors != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("number of bytes written : ", byte1)

	//reading a file
	file1, err1 := os.Open("file.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	// create a buffer to store the content of the file
	buffer := make([]byte, 1024)

	// reading the content of the file into buffer
	for {
		n, err := file1.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

	//readig a file using ioutil
	content1, err2 := ioutil.ReadFile("file.txt") // not suitanle for large files as it reads the whole file at once
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(string(content1))

	content2, err3 := os.ReadFile("file.txt") // read the whole file at once
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	fmt.Println(string(content2))

}
