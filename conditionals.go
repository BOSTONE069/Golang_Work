package main

import "fmt"

func main() {
  if lessonLearned := true; lessonLearned {
    fmt.Println("Great job! You can continue on to the next exercise.")
  } else {
    fmt.Println("Practice makes perfect.")
  }
  // Create more conditions below!
  if ungaPrice := 230; ungaPrice > 100{
    fmt.Println("The economy is bad")
  } else {
    fmt.Println("The economy is doing great")
  }
}

