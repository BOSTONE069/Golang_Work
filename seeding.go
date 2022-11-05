package main

import (
	"fmt"
  "math/rand"
  "time"
)

func main() {
	// Add your code below:
  rand.Seed(time.Now().UnixNano())
  
  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  } else {
    fmt.Println("Where did all my money go?")
  }
}
