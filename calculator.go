package main

import "fmt"


func add(x, y int) {
  return x + y
}

func sub(x, y int) {
  return x - y
}

func mul(x, y int) {
  return x * y
}

func div(x, y int) {
  return x /y
}
  

func main () {
  var action string
  fmt.Println("Welcome to this Golang calculator")
 
  fmt.Scan(&action)

  if action == "add" || action == "ADD" {
    var x, y int
    fmt.Println("Enter the values to be added\n")
    fmt.Scan(&x)
    fmt.Scan(&y)
    add(x, y)
  } 

}

