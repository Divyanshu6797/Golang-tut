package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("current time : ", currentTime)
	fmt.Printf("current time : %T\n", currentTime)

	formattedTime := currentTime.Format("02-01-2006, Monday, 15:04:05, 3:04PM")
	fmt.Println("formatted time : ", formattedTime)

	dateString := "02-01-2006"
	date, _ := time.Parse("02-01-2006", dateString)
	fmt.Println("date : ", date)

	// adding 1 day to current time
	newTime := currentTime.AddDate(0, 0, 1)
	fmt.Println("new time : ", newTime)

}
