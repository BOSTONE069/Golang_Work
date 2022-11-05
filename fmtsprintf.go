package main

import "fmt"

func main() {
  template := "I wish I had a %v."
  pet := "puppy"
  
  var wish string
  
  // Add your code below:
  wish = fmt.Sprintf(template, pet)
  
  fmt.Println(wish)
}
