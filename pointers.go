package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	star := "Polaris"

	// Add your code below:
  starAddress := &star
  fmt.Println("The address of star is", starAddress)
	fmt.Println("The value of star is", *starAddress)
}
