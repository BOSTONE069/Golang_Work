package main

import "fmt"

func main() {
  fmt.Println("What would you like for lunch?")
  
  // Add your code below:
  var food string
  fmt.Scan(&food)
  fmt.Printf("Sure, we can have %v for lunch.", food)
}
