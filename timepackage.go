package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("Welcome to time study of Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//This is time usage with specified default formats in Golang
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(2022, time.November, 24, 4, 34, 0, 0, time.UTC)
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("01-02-2006 Monday"))
}