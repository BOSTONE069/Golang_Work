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

	
}