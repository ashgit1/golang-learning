package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang...")

	// get the current date
	presentTime := time.Now()
	fmt.Println(presentTime)
	// Foramt the current date, the formatting values are by default the same as shown below
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//Create a new Date
	createdDate := time.Date(2021, time.November, 06, 11, 55, 00, 00, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
