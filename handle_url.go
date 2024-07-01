package main

import (
	"fmt"
	"net/url"
)

func main() {
	url1 := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("URL : %T", url1)

	parsedUrl, err := url.Parse(url1)
	if err != nil {
		fmt.Println("cant parse url : ", err)
	}
	fmt.Printf("type of parsedUrl : %T", parsedUrl)
	fmt.Println("Parsed URL : ", parsedUrl)

	fmt.Println("Scheme : ", parsedUrl.Scheme) // indicates the protocol
	fmt.Println("Host : ", parsedUrl.Host)
	fmt.Println("Path : ", parsedUrl.Path)
	fmt.Println("RawQuery : ", parsedUrl.RawQuery)
	fmt.Println("Fragment : ", parsedUrl.Fragment)

	// modifying url components
	parsedUrl.Path = "/newpath"
	parsedUrl.RawQuery = "id=1"
	fmt.Println("Modified URL : ", parsedUrl)

}
