package main

import "fmt"

func main() {
	star := "Polaris"
	
	starAddress := &star
	
	// Add your code below:
	*starAddress = "Sirius"
	
  
  fmt.Println("The actual value of star is", star)
}
