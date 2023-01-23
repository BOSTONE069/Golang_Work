package main

import "fmt"

// The function main() is the entry point of the program
func main() {
  fmt.Println("If else in golang")

  logincount := 23
  var result string

	heistReady := true


  if heistReady {
    fmt.Print("Let's Go!\n")
  }

  if logincount < 10 {
    result = "Regular User"
  } else if logincount > 10 {
    result = "watch out"
  } else {
    result = "Exactly 10 login count"
  }

  fmt.Println(result)
}
